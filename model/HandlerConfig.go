package model

type HandlerConfig struct {
	Op         string     `json:"op"`
	MiddleWare Middleware `json:"middle_ware"`
}

type IdAndHandlerConfig struct {
	Id     string
	Config *HandlerConfig
}
