package model

import "crypto/rsa"

type Cipher interface {
	getServerKeyPair() KeyPair
	getPublicKeyFromClient() []byte
	encrypt(str []byte) []byte
	decrypt(str []byte) []byte
}

type KeyPair struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}
