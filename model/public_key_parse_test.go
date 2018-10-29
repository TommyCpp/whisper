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
	publicKey := []byte(`
-----BEGIN PUBLIC KEY-----
MIIBCgKCAQEAqHQrlA91sjLeLut2+/V7yNJ6KYMM1IFlzQzCG/PMHwM1daf3s32BiHFDQ2Wa
TH5VP4yN6HCgo6QoumwXu35rbY5VtaLBdM7u555BiKQ3iF4OoQ5PkFza+zuskQau5ykhSvsL
/3f+VMXhMRLaI4Am2aL4/+MBOQDLHf21b8cI24OXAaeDNsENUS67rsS5bBOaTUsr6wk/1qjs
HEzAk3VsHlASo6Nr5rHkW3Agff14x0uOh1X5uYJHJDV8Pgz70K0bVsSW14VewZVNnoaoBPzy
76zEiPAbOg6sVYhNFdqbQxXxJJRZ21JLGnQPv39J8X1rbh/oRRXhhhfct3POWEUkMwIDAQAB
-----END PUBLIC KEY-----
`)
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
