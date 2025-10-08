package register

import (
	"fmt"
	"imperishable-gate/internal/client/service"
	"imperishable-gate/internal/client/utils"
)

// HandleRegister 处理用户注册流程（包含邮箱验证）
func HandleRegister(addr string) error {
	// 1. 读取用户输入
	userInput, err := ReadUserInput()
	if err != nil {
		return err
	}

	// 2. 调用注册服务
	fmt.Println("\nSending registration request...")
	err = service.Register(addr, userInput.Username, userInput.Email, userInput.Password)
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}

	fmt.Println("Account created successfully!")
	fmt.Println("A verification code has been sent to your email:", userInput.Email)
	fmt.Println("The code is valid for 15 minutes.")

	// 3. 验证邮箱（使用通用的验证码处理逻辑）
	err = utils.HandleVerificationWithRetry(utils.VerificationConfig{
		MaxAttempts: 5,
		ResendFunc: func() error {
			return service.ResendVerificationEmail(addr, userInput.Email)
		},
		ValidateFunc: func(code string) error {
			return service.VerifyEmail(addr, userInput.Email, code)
		},
		SuccessMessage: "Email verification successful!\nRegistration completed!\nYou can now login with your credentials using 'gate login' command.",
		FailureMessage: "Email verification failed",
	})

	return err
}
