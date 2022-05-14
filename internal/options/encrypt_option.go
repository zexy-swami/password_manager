package options

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/zexy-swami/password_manager/internal/encryption"
	"github.com/zexy-swami/password_manager/pkg/CLI"
)

func encryptOption() error {
	fmt.Printf("\n%s", encFilePathRequest)
	filePath := CLI.InputString()

	fmt.Printf("\n%s", encStoragePathRequest)
	storagePath := CLI.InputString()

	passphrase, err := getPassphrase()
	if err != nil {
		return err
	}

	return encryption.EncryptionProcess(filePath, storagePath, passphrase)
}

func getPassphrase() ([]byte, error) {
	fmt.Printf("\n%s", firstPassphraseRequest)
	firstPassphrase, err := CLI.InputBytesSecure()
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n%s", secondPassphraseRequest)
	secondPassphrase, err := CLI.InputBytesSecure()
	if err != nil {
		return nil, err
	}

	if bytes.Compare(firstPassphrase, secondPassphrase) != 0 {
		return nil, errors.New("mismatched passphrases were entered")
	}

	return firstPassphrase, nil
}
