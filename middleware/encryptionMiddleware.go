package middleware

import "github.com/tommycpp/Whisper/model"

type EncryptionMiddleware struct {
	cipher model.Cipher
}

func (encryption *EncryptionMiddleware) AfterRead(msg string) string {
	return string(encryption.cipher.Decrypt([]byte(msg)))
}

func (encryption *EncryptionMiddleware) BeforeWrite(msg string) string {
	return string(encryption.cipher.Decrypt([]byte(msg)))
}
