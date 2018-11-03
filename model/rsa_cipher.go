package model

import (
	"crypto/rand"
	"crypto/rsa"
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
	//Encrypt with user's public key
	res, err := rsa.EncryptPKCS1v15(rand.Reader, cipher.PublicKeyFromClient, str)
	if err != nil {
		panic("Error when Encrypt")
	} else {
		return res
	}
}

func (cipher *RSACipher) Decrypt(str []byte) []byte {
	res, err := rsa.DecryptPKCS1v15(nil, cipher.KeyPair.PrivateKey, str)
	if err != nil {
		log.Println("fail to Decrypt")
		return nil
	} else {
		return res
	}
}

func NewRSACipher(publicKeyFromClient []byte) *RSACipher {
	block, _ := pem.Decode(publicKeyFromClient)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Println("failed to decode PEM block containing public key")
	}
	var rsaPublicKey *rsa.PublicKey
	rsaPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
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
