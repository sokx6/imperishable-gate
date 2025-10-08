package email

import (
	"fmt"
	"imperishable-gate/internal/server/utils/logger"
	"os"

	"github.com/joho/godotenv"
)

// Config 邮件配置
type Config struct {
	SMTPHost string
	SMTPPort string
	From     string
	Password string
}

// LoadConfig 加载邮件配置
func LoadConfig() (*Config, error) {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		logger.Warning("Failed to load .env file: %v", err)
	}

	// 读取配置
	config := &Config{
		SMTPHost: os.Getenv("EMAIL_HOST"),
		SMTPPort: os.Getenv("EMAIL_PORT"),
		From:     os.Getenv("EMAIL_FROM"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}

	// 验证配置
	if config.SMTPHost == "" || config.SMTPPort == "" || config.From == "" || config.Password == "" {
		return nil, fmt.Errorf("email configuration is incomplete")
	}

	return config, nil
}

// GetSMTPAddress 获取 SMTP 服务器地址
func (c *Config) GetSMTPAddress() string {
	return c.SMTPHost + ":" + c.SMTPPort
}
