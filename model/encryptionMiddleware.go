package model

import "encoding/base64"

type EncryptionMiddleware struct {
	Cipher Cipher
}

type E2eEncryptionMiddleware struct {
	EncryptionMiddleware
}

type RSAEncryptionMiddleware struct {
	EncryptionMiddleware
}

func (encryption *EncryptionMiddleware) AfterRead(msg *Message) error {
	var err error
	var messageBytes []byte
	messageBytes, err = base64.StdEncoding.DecodeString(msg.Content)
	if err != nil {
		return err
	}
	msg.Content = string(encryption.Cipher.Decrypt(messageBytes))
	return nil
}

func (encryption *EncryptionMiddleware) BeforeWrite(msg *Message) error {
	msg.Content = string(encryption.Cipher.Encrypt([]byte(msg.Content)))
	msg.Content = base64.StdEncoding.EncodeToString([]byte(msg.Content))
	return nil
}

func NewEncryptionMiddleware(cipher Cipher) *EncryptionMiddleware {
	return &EncryptionMiddleware{
		Cipher: cipher,
	}
}

func NewRSAEncryptionMiddleware(cipher *RSACipher) *RSAEncryptionMiddleware {
	return &RSAEncryptionMiddleware{
		EncryptionMiddleware: EncryptionMiddleware{
			Cipher: cipher,
		},
	}
}
