package register

import (
	"bufio"
	"fmt"
	"strings"
)

// readUsername 读取并验证用户名
func readUsername(reader *bufio.Scanner) (string, error) {
	fmt.Print("Please enter your username (3-32 characters): ")
	if !reader.Scan() {
		return "", fmt.Errorf("failed to read username")
	}
	username := strings.TrimSpace(reader.Text())
	if username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}
	if len(username) < 3 || len(username) > 32 {
		return "", fmt.Errorf("username must be between 3 and 32 characters")
	}
	return username, nil
}
