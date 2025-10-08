package register

import (
	"bufio"
	"fmt"
)

// readPassword 读取并验证密码
func readPassword(reader *bufio.Scanner) (string, error) {
	fmt.Print("Please enter your password (minimum 6 characters): ")
	if !reader.Scan() {
		return "", fmt.Errorf("failed to read password")
	}
	password := reader.Text()
	if len(password) < 6 {
		return "", fmt.Errorf("password must be at least 6 characters long")
	}
	return password, nil
}
