package model

type Cipher interface {
	getKeyPair() KeyPair
	encryption(str string)
	decryption(str string)
}

type KeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}
