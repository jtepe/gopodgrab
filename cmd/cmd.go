package cmd

import (
	"bufio"
	"fmt"
	"os"
)

// waitApproval blocks until the user approves or denies the progression
// with a "y" or "yes" input.Every other input (or error) is interpreted
// as disapproval. The function adds a (yes/no) substring to the message,
// so don't supply this part yourself
func waitApproval(msg string) bool {
	fmt.Printf(msg + " (yes/no) ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "y" || s.Text() == "yes" {
			return true
		}

		break
	}

	return false
}

// humanized gives a string representation of bytes that is supposed to be
// more understandable for humans.
func humanized(bytes int64) string {
	num := float64(bytes)
	exp := 0

	for num >= 1024 && exp < 4 {
		num = num / 1024
		exp++
	}

	unit := [4]string{"B", "KB", "MB", "GB"}

	return fmt.Sprintf("%.2f %s", num, unit[exp])
}
