// ensure_token.go
package service

import (
	"errors"
	"fmt"

	"imperishable-gate/internal/client/utils"
)

// 这个函数尝试自动获取有效的 access token。
// 如果无法自动获取（例如没有 refresh token），它会提示用户重新登录。
// 如果成功获取到有效的 token，则返回该 token。
// 如果失败，则返回错误。
func EnsureValidTokenWithPrompt(addr, accessToken string) (string, error) {
	token, err := GetTokenAutomatically(addr, accessToken)

	if err != nil {
		switch {
		case errors.Is(err, utils.ErrNoRefreshToken):
			fmt.Println("Authentication expired: no refresh token found.")
			fmt.Println("Please run 'login' to sign in again.")
			return "", err
		default:
			return "", fmt.Errorf("failed to get valid token: %w", err)
		}
	}

	return token, nil
}
