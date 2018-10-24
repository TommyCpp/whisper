package model

import "crypto/rsa"

type Cipher interface {
	Encrypt(str []byte) []byte
	Decrypt(str []byte) []byte
}

type KeyPair struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}
