package decryption

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var outputFormat = "password from %s (login: %s): %s\n"

func DecryptionProcess(filePath string, passphrase []byte) error {
	encryptedData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("error occcurred while reading data from %s", filePath))
	}

	decryptedData, err := dataDecryption(encryptedData, passphrase)
	if err != nil {
		return err
	}

	decryptedDataAsSlice := strings.Split(decryptedData, " ")
	fmt.Printf(outputFormat, decryptedDataAsSlice[0], decryptedDataAsSlice[1], decryptedDataAsSlice[2])

	return nil
}
