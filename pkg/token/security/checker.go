package security

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/ip2region"
	"github.com/infraboard/keyauth/pkg/session"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
)

// NewChecker todo
func NewChecker() (Checker, error) {
	if pkg.Domain == nil {
		return nil, fmt.Errorf("denpence domain service required")
	}
	if pkg.User == nil {
		return nil, fmt.Errorf("denpence user service required")
	}
	if pkg.SessionAdmin == nil {
		return nil, fmt.Errorf("denpence session service required")
	}
	if pkg.IP2Region == nil {
		return nil, fmt.Errorf("denpence ip2region service required")
	}
	c := cache.C()
	if c == nil {
		return nil, fmt.Errorf("denpence cache service is nil")
	}

	return &checker{
		domain:   pkg.Domain,
		user:     pkg.User,
		session:  pkg.SessionAdmin,
		cache:    c,
		ip2Regin: pkg.IP2Region,
		log:      zap.L().Named("Login Security"),
	}, nil
}

type checker struct {
	domain   domain.DomainServiceServer
	user     user.UserServiceServer
	session  session.AdminServiceServer
	cache    cache.Cache
	ip2Regin ip2region.Service
	log      logger.Logger
}

func (c *checker) MaxFailedRetryCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.LoginSecurity.RetryLock {
		c.log.Debugf("retry lock check disabled, don't check")
		return nil
	}
	c.log.Debugf("max failed retry lock check enabled, checking ...")

	var count uint32
	err := c.cache.Get(req.AbnormalUserCheckKey(), count)
	if err != nil {
		c.log.Errorf("get key %s from cache error, %s", req.AbnormalUserCheckKey(), err)
	}

	rc := ss.LoginSecurity.RetryLockConfig
	if count > rc.RetryLimite {
		return fmt.Errorf("retry %d times, reach the max(%d) retry limit", count, rc.RetryLimite)
	}

	return nil
}

func (c *checker) UpdateFailedRetry(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.LoginSecurity.RetryLock {
		c.log.Debugf("retry lock check disabled, don't check")
		return nil
	}

	var count int
	if c.cache.IsExist(req.AbnormalUserCheckKey()) {
		// 之前已经登陆失败过
		err := c.cache.Put(req.AbnormalUserCheckKey(), count+1)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	} else {
		// 首次登陆失败
		err := c.cache.PutWithTTL(
			req.AbnormalUserCheckKey(),
			count+1,
			ss.LoginSecurity.RetryLockConfig.LockedMiniteDuration(),
		)
		if err != nil {
			c.log.Errorf("set key %s to cache error, %s", req.AbnormalUserCheckKey())
		}
	}
	return nil
}

func (c *checker) OtherPlaceLoggedInChecK(ctx context.Context, tk *token.Token) error {
	ss := c.getOrDefaultSecuritySettingWithDomain(tk.Account, tk.Domain)
	if !ss.LoginSecurity.ExceptionLock {
		c.log.Debugf("exception check disabled, don't check")
		return nil
	}

	if !ss.LoginSecurity.ExceptionLockConfig.OtherPlaceLogin {
		c.log.Debugf("other place login check disabled, don't check")
		return nil
	}

	c.log.Debugf("other place login check enabled, checking ...")

	// 查询当前登陆IP地域
	login, err := c.ip2Regin.LookupIP(tk.GetRemoteIP())
	if err != nil {
		return err
	}

	// 查询出用户上次登陆的地域
	queryReq := session.NewQueryUserLastSessionRequest(tk.Account)
	us, err := c.session.QueryUserLastSession(ctx, queryReq)
	if err != nil {
		if exception.IsNotFoundError(err) {
			c.log.Debugf("user %s last login session not found", tk.Account)
			return nil
		}

		return err
	}

	// city为0 表示内网IP, 不错异地登录校验
	if us.IpInfo.CityId == 0 {
		c.log.Warnf("city id is 0, 内网IP skip OtherPlaceLoggedInChecK")
		return nil
	}

	if us != nil {
		c.log.Debugf("user last login city: %s (%d)", us.IpInfo.City, us.IpInfo.CityId)
		if login.CityID != us.IpInfo.CityId {
			return fmt.Errorf("异地登录, 请输入验证码后再次提交")
		}
	}

	return nil
}

func (c *checker) NotLoginDaysChecK(ctx context.Context, tk *token.Token) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, tk.Domain)
	if !ss.LoginSecurity.ExceptionLock {
		c.log.Debugf("exception check disabled, don't check")
		return nil
	}
	c.log.Debugf("not login days check enabled, checking ...")

	// 查询出用户上次登陆的地域
	queryReq := session.NewQueryUserLastSessionRequest(tk.Account)
	us, err := c.session.QueryUserLastSession(ctx, queryReq)
	if err != nil {
		if exception.IsNotFoundError(err) {
			c.log.Debugf("user %s last login session not found", tk.Account)
			return nil
		}

		return err
	}

	if us != nil {
		days := uint32(time.Now().Sub(time.Unix(us.LoginAt/1000, 0)).Hours() / 24)
		c.log.Debugf("user %d days not login", days)
		maxDays := ss.LoginSecurity.ExceptionLockConfig.NotLoginDays
		if days > maxDays {
			return fmt.Errorf("user not login days %d", days)
		}
		c.log.Debugf("not login days check passed, days: %d, max days: %d", days, maxDays)
	}

	return nil
}

func (c *checker) IPProtectCheck(ctx context.Context, req *token.IssueTokenRequest) error {
	ss := c.getOrDefaultSecuritySettingWithUser(ctx, req.Username)
	if !ss.LoginSecurity.IpLimite {
		c.log.Debugf("ip limite check disabled, don't check")
		return nil
	}

	c.log.Debugf("ip limite check enabled, checking ...")

	return nil
}

func (c *checker) getOrDefaultSecuritySettingWithUser(ctx context.Context, account string) *domain.SecuritySetting {
	ss := domain.NewDefaultSecuritySetting()
	u, err := c.user.DescribeAccount(ctx, user.NewDescriptAccountRequestWithAccount(account))
	if err != nil {
		c.log.Errorf("get user account error, %s, use default setting to check", err)
		return ss
	}

	return c.getOrDefaultSecuritySettingWithDomain(u.Data.Profile.Account, u.Domain)
}

func (c *checker) getOrDefaultSecuritySettingWithDomain(account, domainName string) *domain.SecuritySetting {
	ss := domain.NewDefaultSecuritySetting()
	d, err := c.domain.DescribeDomain(pkg.GetInternalAdminTokenCtx(account), domain.NewDescribeDomainRequestWithName(domainName))
	if err != nil {
		c.log.Errorf("get domain error, %s, use default setting to check", err)
		return ss
	}

	return d.SecuritySetting
}
