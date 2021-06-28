// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type GoodsHandler interface {
	CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsReply, error)

	DeleteGoods(context.Context, *DeleteGoodsRequest) (*DeleteGoodsReply, error)

	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error)

	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error)

	UpdateGoods(context.Context, *UpdateGoodsRequest) (*UpdateGoodsReply, error)
}

func NewGoodsHandler(srv GoodsHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/api.goods.service.v1.Goods/CreateGoods", func(w http.ResponseWriter, r *http.Request) {
		var in CreateGoodsRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateGoods(ctx, req.(*CreateGoodsRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*CreateGoodsReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/api.goods.service.v1.Goods/UpdateGoods", func(w http.ResponseWriter, r *http.Request) {
		var in UpdateGoodsRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateGoods(ctx, req.(*UpdateGoodsRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*UpdateGoodsReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/api.goods.service.v1.Goods/DeleteGoods", func(w http.ResponseWriter, r *http.Request) {
		var in DeleteGoodsRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteGoods(ctx, req.(*DeleteGoodsRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*DeleteGoodsReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/goods/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in GetGoodsRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.BindVars(mux.Vars(r), &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGoods(ctx, req.(*GetGoodsRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetGoodsReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/api.goods.service.v1.Goods/ListGoods", func(w http.ResponseWriter, r *http.Request) {
		var in ListGoodsRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListGoods(ctx, req.(*ListGoodsRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ListGoodsReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	return r
}

type GoodsHTTPClient interface {
	CreateGoods(ctx context.Context, req *CreateGoodsRequest, opts ...http1.CallOption) (rsp *CreateGoodsReply, err error)

	DeleteGoods(ctx context.Context, req *DeleteGoodsRequest, opts ...http1.CallOption) (rsp *DeleteGoodsReply, err error)

	GetGoods(ctx context.Context, req *GetGoodsRequest, opts ...http1.CallOption) (rsp *GetGoodsReply, err error)

	ListGoods(ctx context.Context, req *ListGoodsRequest, opts ...http1.CallOption) (rsp *ListGoodsReply, err error)

	UpdateGoods(ctx context.Context, req *UpdateGoodsRequest, opts ...http1.CallOption) (rsp *UpdateGoodsReply, err error)
}

type GoodsHTTPClientImpl struct {
	cc *http1.Client
}

func NewGoodsHTTPClient(client *http1.Client) GoodsHTTPClient {
	return &GoodsHTTPClientImpl{client}
}

func (c *GoodsHTTPClientImpl) CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...http1.CallOption) (out *CreateGoodsReply, err error) {
	path := binding.EncodePath("POST", "/api.goods.service.v1.Goods/CreateGoods", in)
	out = &CreateGoodsReply{}

	err = c.cc.Invoke(ctx, path, nil, &out, http1.Method("POST"), http1.PathPattern("/api.goods.service.v1.Goods/CreateGoods"))

	return
}

func (c *GoodsHTTPClientImpl) DeleteGoods(ctx context.Context, in *DeleteGoodsRequest, opts ...http1.CallOption) (out *DeleteGoodsReply, err error) {
	path := binding.EncodePath("POST", "/api.goods.service.v1.Goods/DeleteGoods", in)
	out = &DeleteGoodsReply{}

	err = c.cc.Invoke(ctx, path, nil, &out, http1.Method("POST"), http1.PathPattern("/api.goods.service.v1.Goods/DeleteGoods"))

	return
}

func (c *GoodsHTTPClientImpl) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...http1.CallOption) (out *GetGoodsReply, err error) {
	path := binding.EncodePath("GET", "/goods/{id}", in)
	out = &GetGoodsReply{}

	err = c.cc.Invoke(ctx, path, nil, &out, http1.Method("GET"), http1.PathPattern("/goods/{id}"))

	return
}

func (c *GoodsHTTPClientImpl) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...http1.CallOption) (out *ListGoodsReply, err error) {
	path := binding.EncodePath("POST", "/api.goods.service.v1.Goods/ListGoods", in)
	out = &ListGoodsReply{}

	err = c.cc.Invoke(ctx, path, nil, &out, http1.Method("POST"), http1.PathPattern("/api.goods.service.v1.Goods/ListGoods"))

	return
}

func (c *GoodsHTTPClientImpl) UpdateGoods(ctx context.Context, in *UpdateGoodsRequest, opts ...http1.CallOption) (out *UpdateGoodsReply, err error) {
	path := binding.EncodePath("POST", "/api.goods.service.v1.Goods/UpdateGoods", in)
	out = &UpdateGoodsReply{}

	err = c.cc.Invoke(ctx, path, nil, &out, http1.Method("POST"), http1.PathPattern("/api.goods.service.v1.Goods/UpdateGoods"))

	return
}
