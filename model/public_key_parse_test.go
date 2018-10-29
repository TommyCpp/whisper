package model

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"testing"
)

func TestParseBase64Key(t *testing.T) {
	publicKey := []byte("-----BEGIN PUBLIC KEY-----\nMIIBCgKCAQEAnS9vYTBp9sHAOlSBzVR3m95/qipHTbMgsWtlM8NKaA7n84nihOiboj0vptle\nWLDxsjytgzFB8ngscjA8wfPpO89FB3mD1EJ03fUOtt1w4gYJLaLLtXRun3DnsEUblQ4VYKD2\n6+n40NhVqKF+F/Nto9sPFa7mcIWH10D6w5SgIiSBUUd3CwGd1ETYr3GcY4yH8SQjQmxRK0Cx\n5OcOX2ag0iiRQ2x3iCRvQrGqWG31xU4DE1DAkgDv8vwFLegB3bmIdkzzE4xmE6FYZ0CW7Xn3\nw+33In8FuP95lPDI/V9svm+68hFLbGI246JOYHY7c/9vkw80SeC7ZvHKWImzbwEEAwIDAQAB\n-----END PUBLIC KEY-----\n")
	//Generate by keypair.js
	block, _ := pem.Decode(publicKey)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	var rsaPublicKey *rsa.PublicKey
	rsaPublicKey = pub
	//if !right {
	//	log.Fatal("Not RSA public key")
	//}
	fmt.Println(rsaPublicKey)
}
