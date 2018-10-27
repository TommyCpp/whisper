package model

type HandlerConfig struct {
	Op         string
	MiddleWare Middleware
}

type IdAndHandlerConfig struct {
	Id     string
	Config *HandlerConfig
}
