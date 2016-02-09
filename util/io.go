package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFullyFromStdin() string {
	input := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = input + line
	}

	return input
}

func AskBool(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")
	text = strings.ToLower(text)

	if err != nil {
		return false
	} else if text == "y" || text == "yes" {
		return true
	} else {
		return false
	}
}
