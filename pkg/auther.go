package pkg

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/infraboard/mcube/exception"
	httpcontext "github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/pkg/domain"
	"github.com/infraboard/keyauth/pkg/endpoint"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user/types"
	"github.com/infraboard/keyauth/version"
)

// GetInternalAdminTokenCtx 内部调用时的模拟token
func GetInternalAdminTokenCtx(account string) context.Context {
	return WithTokenContext(context.Background(), &token.Token{
		Account:  account,
		Domain:   domain.AdminDomainName,
		UserType: types.UserType_INTERNAL,
	})
}

// NewInternalAuther 内部使用的auther
func NewInternalAuther() router.Auther {
	return &internal{}
}

// internal todo
type internal struct {
}

func (i *internal) Auth(r *http.Request, entry router.Entry) (
	authInfo interface{}, err error) {
	var tk *token.Token
	if entry.AuthEnable {
		req := token.NewValidateTokenRequest()
		// 获取需要校验的access token(用户的身份凭证)
		accessToken := r.Header.Get("x-oauth-token")
		if accessToken == "" {
			return nil, exception.NewUnauthorized("x-oauth-token header required")
		}
		req.AccessToken = accessToken

		tk, err = Token.ValidateToken(nil, req)
		if err != nil {
			return nil, err
		}
	}

	if entry.PermissionEnable && tk != nil {
		// 如果是超级管理员不做权限校验, 直接放行
		if tk.UserType.IsIn(types.UserType_SUPPER) {
			return tk, nil
		}

		// 其他比如服务类型, 主账号类型, 子账号类型
		// 如果开启权限认证都需要检查
		if Permission == nil {
			return nil, fmt.Errorf("permission service not load")
		}

		req := permission.NewCheckPermissionrequest()
		req.EnpointID = i.endpointHashID(entry)
		_, err = Permission.CheckPermission(req)
		if err != nil {
			return nil, exception.NewPermissionDeny("no permission")
		}
	}

	return tk, nil
}

func (i *internal) endpointHashID(entry router.Entry) string {
	return endpoint.GenHashID(version.ServiceName, entry.Path, entry.Method)
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// GetTokenFromHTTPRequest 从上下文中获取Token
func GetTokenFromHTTPRequest(r *http.Request) (*token.Token, error) {
	ctx := httpcontext.GetContext(r)

	if ctx.AuthInfo == nil {
		return nil, exception.NewInternalServerError("authInfo is not in request context, please check auth middleware")
	}

	tk, ok := httpcontext.GetContext(r).AuthInfo.(*token.Token)
	if !ok {
		return nil, exception.NewInternalServerError("authInfo is not token pointer")
	}

	tk.WithRemoteIP(request.GetRemoteIP(r))
	tk.WithUerAgent(r.UserAgent())
	return tk, nil
}

// GetTokenCtxFromHTTPRequest todo
func GetTokenCtxFromHTTPRequest(r *http.Request) (context.Context, error) {
	tk, err := GetTokenFromHTTPRequest(r)
	if err != nil {
		return nil, err
	}

	return WithTokenContext(context.Background(), tk), nil
}

// ContextKeyType key类型
type contextKeyType string

const (
	// ContextKeyType_TOKEN token 上下文key
	contextKeyTypeToken = contextKeyType("token")
)

// WithTokenContext todo
func WithTokenContext(ctx context.Context, tk *token.Token) context.Context {
	return context.WithValue(ctx, contextKeyTypeToken, tk)
}

// GetTokenFromContext 从上下文中获取token
func GetTokenFromContext(ctx context.Context) *token.Token {
	if v, ok := ctx.Value(contextKeyTypeToken).(*token.Token); ok {
		return v
	}

	return token.NewDefaultToken()
}
