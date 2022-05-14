package CLI

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func InputString() string {
	CLIReader := bufio.NewReader(os.Stdin)
	receivedString, _ := CLIReader.ReadString('\n')
	formattedString := strings.Map(isGraphicRune, receivedString)
	return formattedString
}

func isGraphicRune(r rune) rune {
	if unicode.IsGraphic(r) {
		return r
	}
	return -1
}
