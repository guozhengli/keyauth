package http

import (
	"net/http"

	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/keyauth/app/endpoint"
)

// CreateApplication 创建自定义角色
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	req := endpoint.NewDefaultRegistryRequest()
	if err := request.GetDataFromRequest(r, req); err != nil {
		response.Failed(w, err)
		return
	}

	_, err := h.endpoint.RegistryEndpoint(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, req)
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	req := endpoint.NewQueryEndpointRequestFromHTTP(r)

	set, err := h.endpoint.QueryEndpoints(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	rctx := context.GetContext(r)

	id := rctx.PS.ByName("id")
	req := endpoint.NewDescribeEndpointRequestWithID(id)

	d, err := h.endpoint.DescribeEndpoint(
		r.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}
