package register

import (
	"fmt"
	"imperishable-gate/internal/client/service"
)

// HandleRegister 处理用户注册流程（包含邮箱验证）
func HandleRegister(addr string) error {
	// 1. 读取用户输入
	userInput, err := ReadUserInput()
	if err != nil {
		return err
	}

	// 2. 调用注册服务
	fmt.Println("\n📤 Sending registration request...")
	err = service.Register(addr, userInput.Username, userInput.Email, userInput.Password)
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}

	fmt.Println("✓ Account created successfully!")
	fmt.Println("📧 A verification code has been sent to your email:", userInput.Email)
	fmt.Println("💡 The code is valid for 15 minutes.")

	// 3. 验证邮箱（支持重试）
	maxAttempts := 5
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		// 读取验证码
		verificationCode, err := ReadVerificationCode(attempt, maxAttempts)
		if err != nil {
			return err
		}

		// 特殊命令：重新发送验证码
		if verificationCode == "resend" {
			fmt.Println("\n📧 Resending verification email...")
			if err := service.ResendVerificationEmail(addr, userInput.Email); err != nil {
				fmt.Println("❌ Failed to resend verification email:", err)
				continue
			}
			fmt.Println("✓ Verification email has been resent!")
			attempt-- // 不计入尝试次数
			continue
		}

		// 验证邮箱
		fmt.Println("\n Verifying your email...")
		err = service.VerifyEmail(addr, userInput.Email, verificationCode)
		if err != nil {
			if attempt < maxAttempts {
				fmt.Printf(" Verification failed: %v\n", err)
				fmt.Printf(" You have %d attempt(s) remaining.\n", maxAttempts-attempt)
				fmt.Println(" Tip: Enter 'resend' to get a new verification code.")
				continue
			}
			return fmt.Errorf("email verification failed after %d attempts: %w", maxAttempts, err)
		}

		// 验证成功
		fmt.Println("\n Email verification successful!")
		fmt.Println("✓ Registration completed!")
		fmt.Println(" You can now login with your credentials using 'gate login' command.")
		return nil
	}

	return fmt.Errorf("verification failed: maximum attempts exceeded")
}
