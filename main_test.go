package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/tommycpp/Whisper/model"
	"log"
	"os"
	"strings"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	res := generateToken("testUser")
	fmt.Println(string(res))
}

func TestGetHandlerConfig(t *testing.T) {
	jsonStr := `{
  "op": "ADD",
  "middleware_name": "RSA",
  "setting": {
    "public_key": "-----BEGIN PUBLIC KEY-----\nMIIBCgKCAQEAnS9vYTBp9sHAOlSBzVR3m95/qipHTbMgsWtlM8NKaA7n84nihOiboj0vptle\nWLDxsjytgzFB8ngscjA8wfPpO89FB3mD1EJ03fUOtt1w4gYJLaLLtXRun3DnsEUblQ4VYKD2\n6+n40NhVqKF+F/Nto9sPFa7mcIWH10D6w5SgIiSBUUd3CwGd1ETYr3GcY4yH8SQjQmxRK0Cx\n5OcOX2ag0iiRQ2x3iCRvQrGqWG31xU4DE1DAkgDv8vwFLegB3bmIdkzzE4xmE6FYZ0CW7Xn3\nw+33In8FuP95lPDI/V9svm+68hFLbGI246JOYHY7c/9vkw80SeC7ZvHKWImzbwEEAwIDAQAB\n-----END PUBLIC KEY-----\n"
  }
}
`
	var handlerConfigString = new(struct {
		Op             string                      `json:"op"`
		MiddlewareName string                      `json:"middleware_name"`
		Settings       map[string]*json.RawMessage `json:"setting"`
	})
	err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(handlerConfigString)
	if err != nil {
		log.Println(err)
	} else {
		switch handlerConfigString.MiddlewareName {
		case "RSA":
			{
				// Add a RSA Middleware
				var publicKey string
				err = json.Unmarshal(*handlerConfigString.Settings["public_key"], &publicKey)
				if err != nil {
					log.Println("Do not have public_key")
				} else {
					middleWare := model.NewRSAEncryptionMiddleware(model.NewRSACipher([]byte(publicKey)))
					fmt.Println(middleWare)
				}
			}

		}
	}

}

func TestPassPublicKey(t *testing.T) {
	reader := rand.Reader
	bitSize := 1024

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		panic("Error when generate key")
	}
	publicKey := &key.PublicKey
	derPkix := x509.MarshalPKCS1PublicKey(publicKey)
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	err = pem.Encode(os.Stdout, block)
	if err != nil {
		log.Println("Error when encode pem")
	}
}
