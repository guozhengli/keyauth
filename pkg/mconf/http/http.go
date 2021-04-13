package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/mconf"
)

var (
	api = &handler{}
)

type handler struct {
	service mconf.MicroConfigServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("config_group")
	r.BasePath("config_group")
	r.Handle("GET", "/", h.QueryGroup)
	r.Handle("POST", "/", h.CreateGroup)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Mconf()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("service_config", api)
}
