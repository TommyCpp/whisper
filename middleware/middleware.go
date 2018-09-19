package middleware

type Middleware interface {
	ReadMiddleware
	WriteMiddleware
}

type ReadMiddleware interface {
	afterRead(msg string)
}

type WriteMiddleware interface {
	beforeWrite(msg string)
}
