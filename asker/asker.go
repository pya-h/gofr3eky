package asker

import (
	"bufio"
	"fmt"
	"os"
)

func Ask(messages ...string) ([]string, []error) {
	input_count := len(messages)
	answers := make([]string, input_count)
	errors := make([]error, input_count)
	reader := bufio.NewReader(os.Stdin)

	for i := range messages {
		fmt.Print(messages[i])

		answers[i], errors[i] = reader.ReadString('\n')
	}
	return answers, errors
}
