package CLI

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func InputBytesSecure() ([]byte, error) {
	byteData, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}
	fmt.Println()

	return byteData, nil
}
