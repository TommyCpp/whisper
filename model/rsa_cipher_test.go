package model

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCipher(t *testing.T) {
	publicKey := []byte(`
-----BEGIN PUBLIC KEY-----
MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgEbZKNDZRLSgISZbmR8lK13qge4Z
NqgdH+tTKHmJuviJmETiQMaaYscHXQx8nCBIvaA4bwuvMM4i2PLea+7cBHIPWO8T
OmfRmssyadxPKxbsc9bK5QsroeQ1fMOWkuOK3g+pIv4KkoBgPfdKuFKLSn7RhHam
YAqi002vBnEBRsefAgMBAAE=
-----END PUBLIC KEY-----`)
	cipher := NewRSACipher(publicKey)
	fmt.Println(cipher)
}

func TestEncryptAndDecrypt(t *testing.T) {
	const clientPub = `
-----BEGIN PUBLIC KEY-----
MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgEbZKNDZRLSgISZbmR8lK13qge4Z
NqgdH+tTKHmJuviJmETiQMaaYscHXQx8nCBIvaA4bwuvMM4i2PLea+7cBHIPWO8T
OmfRmssyadxPKxbsc9bK5QsroeQ1fMOWkuOK3g+pIv4KkoBgPfdKuFKLSn7RhHam
YAqi002vBnEBRsefAgMBAAE=
-----END PUBLIC KEY-----
`
	const clientPri = `
-----BEGIN RSA PRIVATE KEY-----
MIICWgIBAAKBgEbZKNDZRLSgISZbmR8lK13qge4ZNqgdH+tTKHmJuviJmETiQMaa
YscHXQx8nCBIvaA4bwuvMM4i2PLea+7cBHIPWO8TOmfRmssyadxPKxbsc9bK5Qsr
oeQ1fMOWkuOK3g+pIv4KkoBgPfdKuFKLSn7RhHamYAqi002vBnEBRsefAgMBAAEC
gYAj39hEEJAyqha/FoitdaPE9Xb/OnMrozvDbCNFj5FGQl4BG1PTfN9hin/6T6q6
yjqCw7CvCPG8n3adXDTpCS2Sj1oitGpONuQwjg+UH8qGvUwc5RFltAE9M+FwMlKT
9MAxFeLYTbaHcv5qrF/EaNfGAdh8VAp0wdOk87BJJVEvoQJBAI06JtZ5hqIxXj93
k3H3PgZ01kafMajOQkMrjhLPHROu9V8w22dX8Kerpfo7Xlgk36Voz4D2Od+jLR1V
Y2D7TbECQQCAbOmW/t37c+P2t7s/0Ub1ydlLr51n1a7aCQh+6A24ZF8TsQZgL1AK
+h/MOPgdLNX4WUn/0LvgqU48QOtFBC5PAkAw7XGhIm8rZ/EgCdxSQncBo57MzsBU
nEjGnqNVDt4jAJ1Pwkxw7D2ayVPycnkIDpZQ5xPkuOlp+k1Z+Ug5xDaxAkAZR9dP
zwoZpr2YYqCstmC2n65z1LUyrIDIEQEoIjwZMUD6Gl377zRdhNFfnVNSQvI3+jOz
9P4XAp0RBWKK6oDHAkB1Z7KhsW9Bng05MtWUBdWbJi5eWnQ4bJ7zKu1BO/cXVHZu
5ilaYh6XOP6hASb+qz6Be7GmU3JMNvopXqyFaVZb
-----END RSA PRIVATE KEY-----
`

	cipher := NewRSACipher([]byte(clientPub))

	encryptedMessage := cipher.Encrypt([]byte("Classification words"))
	fmt.Println(encryptedMessage)

	wordsFromClient, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, cipher.KeyPair.PublicKey, []byte("Classification from client"), nil)

	decryptedMessage := cipher.Decrypt(wordsFromClient)

	assert.Equal(t, "Classification from client", string(decryptedMessage))
}
