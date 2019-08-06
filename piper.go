package piper

import (
	"bufio"
	"os"
	"strings"
)

// Read gets input from stdin, and returns it as a string.
func Read() (string, error) {
	var err error
	var lines []string
	var out string

	// Check for piped in input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Body input is being piped in
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err = scanner.Err(); err != nil {
			return "", err
		}
		out = strings.Join(lines, "\n")
	}

	return strings.Trim(out, "\n"), nil
}
