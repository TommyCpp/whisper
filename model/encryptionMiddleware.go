package model

type EncryptionMiddleware struct {
	Cipher Cipher
}

type E2eEncryptionMiddleware struct {
	EncryptionMiddleware
}

func (encryption *EncryptionMiddleware) AfterRead(msg *Message) error {
	msg.Content = string(encryption.Cipher.Decrypt([]byte(msg.Content)))
	return nil
}

func (encryption *EncryptionMiddleware) BeforeWrite(msg *Message) error {
	msg.Content = string(encryption.Cipher.Decrypt([]byte(msg.Content)))
	return nil
}

func NewEncryptionMiddleware(cipher Cipher) *EncryptionMiddleware {
	return &EncryptionMiddleware{
		Cipher: cipher,
	}
}
