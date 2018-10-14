package model

type RSACipher struct {
	KeyPair KeyPair
}

func (cipher *RSACipher) getKeyPair() KeyPair {
	panic("implement me")
}

func (cipher *RSACipher) encryption(str string) {

}

func (cipher *RSACipher) decryption(str string) {

}

func NewCipher() *RSACipher {
	return &RSACipher{
		KeyPair: generateKeys(),
	}
}

func generateKeys() KeyPair {

}
