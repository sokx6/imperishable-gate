package register

import (
	"fmt"
	"imperishable-gate/internal/client/service"
)

// HandleRegister 处理用户注册流程
func HandleRegister(addr string) error {
	// 读取用户输入
	userInput, err := ReadUserInput()
	if err != nil {
		return err
	}

	// 调用注册服务
	err = service.Register(addr, userInput.Username, userInput.Email, userInput.Password)
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}

	fmt.Println("\n✓ Registration successful!")
	fmt.Println("You can now login with your credentials using 'gate login' command.")
	return nil
}
