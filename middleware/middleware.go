package middleware

type Middleware interface {
	ReadMiddleware
	WriteMiddleware
}

type ReadMiddleware interface {
	AfterRead(msg string) string
}

type WriteMiddleware interface {
	BeforeWrite(msg string) string
}
