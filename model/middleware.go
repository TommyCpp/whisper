package model

type Middleware interface {
	ReadMiddleware
	WriteMiddleware
}

type ReadMiddleware interface {
	AfterRead(msg *Message) error
}

type WriteMiddleware interface {
	BeforeWrite(msg *Message) error
}
