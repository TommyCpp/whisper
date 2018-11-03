package model

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
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
MIIBCgKCAQEAuOTGYgAT9A4cUJy5XK/U+NbFZ5ZE6oB7ygKLpp/NMGVshzH4N35lzuOSeuET
Cqh2qqRefujzSOdT6DzSPUzvsXiz3Aa0FOO8vT93zQMsmsdSnvUMLxAFPYJO/sEnK45u6O6Z
EdMARfd0jSEiEVpVKpUtnbE4pODw42ZSsI8faxYnbeZUnlaw+5VUNq29howuBWMcHKJ+F0tH
5hnHGrl39h9uxpMQ1iYVCxvWbj3dxSI9dGj3vYwOeuYa7BQq4vwm7qcKgezo+utpr4ShHJhY
gRmFDx09hrQkwqHKFScMQwDlng9MY+VwFtqAWcTe9u+nf0aaKwliRoER1PU5VoESjwIDAQAB
-----END PUBLIC KEY-----
`
	cipher := NewRSACipher([]byte(clientPub))

	encryptedMessage := cipher.Encrypt([]byte("Classification words"))
	fmt.Println(encryptedMessage)

	wordsFromClient, _ := rsa.EncryptPKCS1v15(rand.Reader, cipher.KeyPair.PublicKey, []byte("Classification from client"))

	decryptedMessage := cipher.Decrypt(wordsFromClient)

	assert.Equal(t, "Classification from client", string(decryptedMessage))
}

func TestRSACipher_Encrypt(t *testing.T) {
	const clientPub = `
-----BEGIN PUBLIC KEY-----
MIIBCgKCAQEAuOTGYgAT9A4cUJy5XK/U+NbFZ5ZE6oB7ygKLpp/NMGVshzH4N35lzuOSeuET
Cqh2qqRefujzSOdT6DzSPUzvsXiz3Aa0FOO8vT93zQMsmsdSnvUMLxAFPYJO/sEnK45u6O6Z
EdMARfd0jSEiEVpVKpUtnbE4pODw42ZSsI8faxYnbeZUnlaw+5VUNq29howuBWMcHKJ+F0tH
5hnHGrl39h9uxpMQ1iYVCxvWbj3dxSI9dGj3vYwOeuYa7BQq4vwm7qcKgezo+utpr4ShHJhY
gRmFDx09hrQkwqHKFScMQwDlng9MY+VwFtqAWcTe9u+nf0aaKwliRoER1PU5VoESjwIDAQAB
-----END PUBLIC KEY-----
`
	const pub = `
-----BEGIN PUBLIC KEY-----
MIGJAoGBAOCs7gYeB0ZX44VKlE6N64yNrCGOG+u/Bh/515CURHrQAD8m4Ojd1wZw
7otlan6iiws3BluWUXlK5sEpSrIjpQR7wxxCySbOwqqw5fAYhjD0FM2x+O694LGO
v2uQGkX2iGt5Z4uCanOH5+cI4UPZm5UsedfDMk+MtOyT17typ9WNAgMBAAE=
-----END PUBLIC KEY-----
`
	const pri = `
-----BEGIN PRIVATE KEY-----
MIICXAIBAAKBgQDgrO4GHgdGV+OFSpROjeuMjawhjhvrvwYf+deQlER60AA/JuDo
3dcGcO6LZWp+oosLNwZbllF5SubBKUqyI6UEe8McQskmzsKqsOXwGIYw9BTNsfju
veCxjr9rkBpF9ohreWeLgmpzh+fnCOFD2ZuVLHnXwzJPjLTsk9e7cqfVjQIDAQAB
AoGAcE0/9ILR9BE+QoPSuakqkejGn0cfIakr8JO7ciMKT7DkTqyqQvuP3UJZmgep
QX8RrRtl7CWot83+pZJ0KbKzaihtVfD5NGn2ywSornI9Vo6qwHw7XZnSJRkoCQqT
sceAEmULLV0bTNF9aHdhc5zSj2jbh5KX59+8UCRmeQgMxUECQQDop+BcWgIP+chs
bk7wtIqzxWigSCCKSYcP7Rl7sxuoBiMpXdAxA1Ewt8XaujxoiolSf0cC7hm+mxKB
zn2aOLOVAkEA9zgRan8rOPQDrCxVEAg468Iruz5ZmBuG/3IaKs9xvbNEWX9TsL1B
OAU96tHuzpgB7l5Hrk/Zx6D43wQGkxIcGQJAD7SFaLaKvRlXdjpcCdOmKUyCK4+y
4qLkAyc2OSt2CnmflgNHMofOy0MckA9SVJxFeNQurvvzsPI25ZxSzj5VoQJAczMK
UpD9yCVVDMb/wF/Efn/VtwQf5dR1/NTjwq01+Erv/7BohEQ8fulaZ/D5kgWdaMFA
L8b/2Zl2Px32HlRjCQJBAJ5cFniYp0Mvx7yHA1TRQBZsqhLp5YdmKBk3TRmo7zr5
OElXChxGA4WLR+mL/hkpjr4N7Skmv9aOWmQ6i58BHyk=
-----END PRIVATE KEY-----
`
	cipher := NewRSACipher([]byte(clientPub))
	const s = "This is a encrypted message"
	//var privateKey = &pem.Block{
	//	Type:  "PRIVATE KEY",
	//	Bytes: x509.MarshalPKCS1PrivateKey(cipher.KeyPair.PrivateKey),
	//}
	//var publicKey = &pem.Block{
	//	Type:  "PUBLIC KEY",
	//	Bytes: x509.MarshalPKCS1PublicKey(cipher.KeyPair.PublicKey),
	//}
	//fmt.Println(pem.Encode(os.Stdout, publicKey))
	//fmt.Println(pem.Encode(os.Stdout, privateKey))
	var err error
	pubBlock, _ := pem.Decode([]byte(pub))
	priBlock, _ := pem.Decode([]byte(pri))
	if cipher.KeyPair.PublicKey, err = x509.ParsePKCS1PublicKey(pubBlock.Bytes); err != nil {
		log.Println(err)
	}
	if cipher.KeyPair.PrivateKey, err = x509.ParsePKCS1PrivateKey(priBlock.Bytes); err != nil {
		log.Println(err)
	}
	var afterEncrypt = `HJXLz10UcEV1qNFzr55qRAz1qQ8xewxx0o4OVy7Hy9iwhGrjfOdQdWDv+fJ/TMLPS86D/4QwLUnPSZRBPDugllLXJjQElaCc6u4VDEyczrgluA5akDbuWOKzwBg+qJ9YaK3Yt+qQvNl5zHC/5tbq2XX0D5Z4Wx1C6rKAug/9ap0=`
	decode64, err := base64.StdEncoding.DecodeString(afterEncrypt)
	if err != nil {
		log.Println(err)
	}
	var result = string(cipher.Decrypt(decode64))
	assert.True(t, result == s)

}
