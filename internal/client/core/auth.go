package core

import (
	"imperishable-gate/internal/client/service"

	"github.com/spf13/cobra"
)

// CommandsWithoutAuth 定义不需要认证的命令列表
var CommandsWithoutAuth = []string{
	"login",
	"register",
	"logout",
	"help",
	"version",
	"ping",
	"pwd",
	"gate",
}

// ShouldSkipAuth 判断命令是否应该跳过认证
func ShouldSkipAuth(cmd *cobra.Command) bool {
	cmdName := cmd.Name()
	cmdUse := cmd.Use

	for _, name := range CommandsWithoutAuth {
		if cmdName == name || cmdUse == name {
			return true
		}
	}
	return false
}

// EnsureAuthentication 确保用户已认证（如果需要）
func EnsureAuthentication(cmd *cobra.Command, addr string, accessToken *string) error {
	if ShouldSkipAuth(cmd) {
		return nil
	}

	var err error
	*accessToken, err = service.EnsureValidTokenWithPrompt(addr, *accessToken)
	return err
}
