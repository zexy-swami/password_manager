package options

import (
	"fmt"

	"github.com/zexy-swami/password_manager/internal/decryption"
	"github.com/zexy-swami/password_manager/pkg/CLI"
)

func decryptOption() error {
	fmt.Printf("\n%s", decFilePathRequest)
	filePath := CLI.InputString()

	fmt.Printf("\n%s", decPassphraseRequest)
	passphrase, err := CLI.InputBytesSecure()
	if err != nil {
		return err
	}

	return decryption.DecryptionProcess(filePath, passphrase)
}
