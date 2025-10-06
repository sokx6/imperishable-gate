package service

import (
	"imperishable-gate/internal/client/utils"
)

func GetTokenAutomatically(addr string, accessToken string) (string, error) {
	// 检查 access token 是否过期
	isExpired, err := utils.IsTokenExpired(accessToken)
	// 检查过程中出错
	if err != nil {
		return "", err
	}
	// 过期则尝试刷新
	if isExpired {
		// 从系统密钥链获取 refresh token
		refreshToken, err := utils.LoadRefreshToken()
		// 获取过程中出错
		if err != nil {
			return "", err
		}
		// 没有 refresh token，无法刷新
		if refreshToken == "" {
			return "", utils.ErrNoRefreshToken
		}
		// 使用 refresh token 刷新 access token
		newToken, err := utils.RefreshToken(refreshToken, addr)
		// 刷新过程中出错
		if err != nil {
			return "", err
		}
		// 刷新成功，更新全局 access token
		return newToken, nil
	}
	return accessToken, nil
}
