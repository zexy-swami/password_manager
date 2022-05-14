package encryption

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func EncryptionProcess(filePath, storagePath string, passphrase []byte) error {
	sourceFD, err := os.Open(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("file %s not found", filePath))
	}
	defer sourceFD.Close()

	sourceScanner := bufio.NewScanner(sourceFD)
	for sourceScanner.Scan() {
		dataToEncrypt := sourceScanner.Bytes()
		encryptedData := dataEncryption(dataToEncrypt, passphrase)

		encryptedFilename := strings.Split(string(dataToEncrypt), " ")[0]
		encryptedFilePath := filepath.Join(storagePath, encryptedFilename)
		outFD, err := os.Create(encryptedFilePath)
		if err != nil {
			return errors.New(fmt.Sprintf("can't create file %s", encryptedFilePath))
		}

		outFD.Write(encryptedData)
		outFD.Close()
	}

	return nil
}
