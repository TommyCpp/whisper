package model

import "encoding/json"

type HandlerConfig struct {
	Op         string     `json:"op"`
	MiddleWare Middleware `json:"middle_ware"`
}

type IdAndHandlerConfig struct {
	Id     string
	Config *HandlerConfig
}

type HandlerConfigJson struct {
	Op             string                      `json:"op"`
	MiddlewareName string                      `json:"middleware_name"`
	Settings       map[string]*json.RawMessage `json:"setting"`
}
