package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadLine 读取一行输入
func ReadLine(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewScanner(os.Stdin)
	if !reader.Scan() {
		return "", fmt.Errorf("failed to read input")
	}
	return strings.TrimSpace(reader.Text()), nil
}

// ConfirmInput 确认两次输入一致
func ConfirmInput(prompt string, origin string) error {
	input, err := ReadLine(prompt)
	if err != nil {
		return err
	}
	if input != origin {
		return fmt.Errorf("inputs do not match")
	}
	return nil
}
