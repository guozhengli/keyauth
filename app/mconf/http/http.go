package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/mconf"
)

var (
	api = &handler{}
)

type handler struct {
	service mconf.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("config_group")
	r.BasePath("config_group")
	r.Handle("GET", "/", h.QueryGroup)
	r.Handle("POST", "/", h.CreateGroup)
	r.Handle("DELETE", "/:name", h.DestroyGroup)

	r.BasePath("config_item")
	r.Handle("GET", "/group/:name", h.QueryItem)
	r.Handle("POST", "/group/:name", h.AddItemToGroup)
}

func (h *handler) Config() error {
	h.service = app.GetGrpcApp(mconf.AppName).(mconf.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return mconf.AppName
}

func init() {
	app.RegistryHttpApp(api)
}
