package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/zexy-swami/password_manager/pkg/digest"
)

func dataEncryption(dataToEncrypt, passphrase []byte) []byte {
	aesKeyAsDigest := digest.GenerateDigest(passphrase)
	aesBlock, _ := aes.NewCipher(aesKeyAsDigest)
	gcm, _ := cipher.NewGCM(aesBlock)

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	return gcm.Seal(nonce, nonce, dataToEncrypt, nil)
}
