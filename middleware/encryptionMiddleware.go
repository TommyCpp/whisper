package middleware

type EncryptionMiddleware struct {
}

func (encryption *EncryptionMiddleware) afterRead(msg string) {
}

func (encryption *EncryptionMiddleware) beforeWrite(msg string) {
}
