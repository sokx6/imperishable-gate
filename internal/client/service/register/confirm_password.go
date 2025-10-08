package register

import (
	"bufio"
	"fmt"
)

// confirmPassword 确认密码
func confirmPassword(reader *bufio.Scanner, password string) error {
	fmt.Print("Please confirm your password: ")
	if !reader.Scan() {
		return fmt.Errorf("failed to read password confirmation")
	}
	confirmPassword := reader.Text()
	if password != confirmPassword {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}
