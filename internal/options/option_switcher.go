package options

import (
	"errors"
	"fmt"

	"github.com/zexy-swami/password_manager/pkg/CLI"
)

func OptionSwitcher() (err error) {
	fmt.Printf("\n%s", optionRequest)
	requestedOption := CLI.InputString()

	switch requestedOption {
	case "e":
		err = encryptOption()
	case "d":
		err = decryptOption()
	default:
		err = errors.New("unidentified option request")
	}

	return
}
