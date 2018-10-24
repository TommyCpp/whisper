package model

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
)

type RSACipher struct {
	KeyPair             *KeyPair
	PublicKeyFromClient *rsa.PublicKey
}

func (cipher *RSACipher) getServerKeyPair() *KeyPair {
	return cipher.KeyPair
}

func (cipher *RSACipher) getPublicKeyFromClient() *rsa.PublicKey {
	return cipher.PublicKeyFromClient
}

func (cipher *RSACipher) Encrypt(str []byte) []byte {
	res, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, cipher.PublicKeyFromClient, str, nil)
	if err != nil {
		panic("Error when Encrypt")
	} else {
		return res
	}
}

func (cipher *RSACipher) Decrypt(str []byte) []byte {
	res, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, cipher.KeyPair.PrivateKey, str, nil)
	if err != nil {
		log.Fatal("fail to Decrypt")
	}
	return res
}

func NewRSACipher(publicKeyFromClient []byte) *RSACipher {
	block, _ := pem.Decode(publicKeyFromClient)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	var rsaPublicKey *rsa.PublicKey
	rsaPublicKey, right := pub.(*rsa.PublicKey)
	if !right {
		log.Fatal("Not RSA public key")
	}
	return &RSACipher{
		KeyPair:             generateKeys(),
		PublicKeyFromClient: rsaPublicKey,
	}
}

func generateKeys() *KeyPair {
	reader := rand.Reader
	bitSize := 1024

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		panic("Error when generate key")
	}
	return &KeyPair{
		PrivateKey: key,
		PublicKey:  &key.PublicKey,
	}

}
