package model

type E2eCipher struct {
	targetId []byte //Id of the target of the other end
}

func (*E2eCipher) Encrypt(str []byte) []byte {
	return str
}

func (*E2eCipher) Decrypt(str []byte) []byte {
	return str
}

func NewE2eCipher(targetId []byte) *E2eCipher {
	return &E2eCipher{
		targetId: targetId,
	}
}
