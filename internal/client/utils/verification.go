package utils

import (
	"fmt"
	"strings"
)

const (
	// VerificationCodeLength 验证码长度
	VerificationCodeLength = 6
	// DefaultMaxAttempts 默认最大尝试次数
	DefaultMaxAttempts = 5
)

// VerificationConfig 验证码处理配置
type VerificationConfig struct {
	MaxAttempts    int
	ResendFunc     func() error            // 重发验证码函数
	ValidateFunc   func(code string) error // 验证验证码函数
	SuccessMessage string
	FailureMessage string
}

// ReadVerificationCode 读取验证码（支持 resend 命令）
func ReadVerificationCode(attempt, maxAttempts int) (string, error) {
	prompt := "Please enter the verification code (or 'resend' to get a new code): "
	if attempt > 1 {
		fmt.Printf("\n[Attempt %d/%d] ", attempt, maxAttempts)
	}

	code, err := ReadLine(prompt)
	if err != nil {
		return "", err
	}

	// 允许输入 "resend" 来重新发送验证码
	if strings.ToLower(code) == "resend" {
		return "resend", nil
	}

	return code, nil
}

// HandleVerificationWithRetry 处理验证码验证流程（带重试机制）
func HandleVerificationWithRetry(config VerificationConfig) error {
	if config.MaxAttempts <= 0 {
		config.MaxAttempts = DefaultMaxAttempts
	}

	for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
		// 读取验证码
		code, err := ReadVerificationCode(attempt, config.MaxAttempts)
		if err != nil {
			return err
		}

		// 处理 resend 命令
		if code == "resend" {
			fmt.Println("\nResending verification code...")
			if err := config.ResendFunc(); err != nil {
				fmt.Println("Failed to resend:", err)
				attempt-- // 不计入尝试次数
				continue
			}
			fmt.Println("Verification code has been resent!")
			attempt-- // 不计入尝试次数
			continue
		}

		// 验证验证码格式
		if err := ValidateVerificationCode(code); err != nil {
			fmt.Println(err)
			continue
		}

		// 执行验证
		fmt.Println("\nVerifying...")
		if err := config.ValidateFunc(code); err != nil {
			if attempt < config.MaxAttempts {
				fmt.Printf("%s: %v\n", config.FailureMessage, err)
				fmt.Printf("You have %d attempt(s) remaining.\n", config.MaxAttempts-attempt)
				fmt.Println("Tip: Enter 'resend' to get a new verification code.")
				continue
			}
			return fmt.Errorf("%s after %d attempts: %w", config.FailureMessage, config.MaxAttempts, err)
		}

		// 验证成功
		fmt.Println("\n", config.SuccessMessage)
		return nil
	}

	return fmt.Errorf("verification failed: maximum attempts exceeded")
}

// ValidateVerificationCode 验证验证码格式
func ValidateVerificationCode(code string) error {
	if code == "" {
		return fmt.Errorf("verification code cannot be empty")
	}
	if len(code) != VerificationCodeLength {
		return fmt.Errorf("verification code must be 6 digits")
	}
	return nil
}

// ValidateUsername 验证用户名格式
func ValidateUsername(username string) error {
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	if len(username) < MinUsernameLength || len(username) > MaxUsernameLength {
		return fmt.Errorf("username must be between %d and %d characters", MinUsernameLength, MaxUsernameLength)
	}
	return nil
}

// ValidateEmail 简单邮箱格式校验
func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

// ValidatePassword 密码长度校验
func ValidatePassword(password string, minLen int) error {
	if len(password) < minLen {
		return fmt.Errorf("password must be at least %d characters long", minLen)
	}
	return nil
}
