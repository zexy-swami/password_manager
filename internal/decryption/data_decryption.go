package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"github.com/zexy-swami/password_manager/pkg/digest"
)

func dataDecryption(dataToDecrypt, passphrase []byte) (string, error) {
	aesKeyAsDigest := digest.GenerateDigest(passphrase)
	aesBlock, _ := aes.NewCipher(aesKeyAsDigest)
	gcm, _ := cipher.NewGCM(aesBlock)

	nonceSize := gcm.NonceSize()
	nonce, encryptedData := dataToDecrypt[:nonceSize], dataToDecrypt[nonceSize:]

	decryptedData, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return "", errors.New("error occurred while decryption process")
	}

	return string(decryptedData), nil
}
