package main

import (
	"github.com/fatih/color"
	"github.com/zexy-swami/password_manager/internal/options"
)

func main() {
	if err := options.OptionSwitcher(); err != nil {
		color.Red(err.Error())
	}
}
