package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// InteractiveModeHandler 交互模式处理器
type InteractiveModeHandler struct {
	RootCmd           *cobra.Command
	InInteractiveMode *bool
}

// Run 启动交互式命令行模式
func (h *InteractiveModeHandler) Run() {
	fmt.Println("Welcome to gate interactive mode")
	fmt.Println("Type commands to execute, 'exit' or 'quit' to exit")
	fmt.Println("Type 'help' to see available commands")
	fmt.Println()

	// 设置交互模式标志
	*h.InInteractiveMode = true
	defer func() { *h.InInteractiveMode = false }()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("gate> ")

		// 读取用户输入
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		// 跳过空行
		if line == "" {
			continue
		}

		// 检查退出命令
		if line == "exit" || line == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		// 解析输入为命令和参数
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		// 特殊处理 help 命令
		if parts[0] == "help" || parts[0] == "--help" || parts[0] == "-h" {
			h.handleHelpCommand(parts)
			fmt.Println()
			continue
		}

		// 特殊处理 version 命令
		if parts[0] == "version" || parts[0] == "--version" || parts[0] == "-v" {
			fmt.Printf("gate version %s\n", h.RootCmd.Version)
			fmt.Println()
			continue
		}

		// 执行命令
		h.executeCommand(parts)

		fmt.Println() // 在命令之间添加空行
	}

	// 检查扫描器错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

// handleHelpCommand 处理 help 命令
func (h *InteractiveModeHandler) handleHelpCommand(parts []string) {
	if len(parts) == 1 {
		// 显示所有可用命令
		h.RootCmd.Help()
	} else {
		// 显示特定命令的帮助
		helpCmd := parts[1]
		found := false
		for _, c := range h.RootCmd.Commands() {
			if c.Name() == helpCmd || c.HasAlias(helpCmd) {
				c.Help()
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Unknown command: %s\n", helpCmd)
		}
	}
}

// executeCommand 执行命令
func (h *InteractiveModeHandler) executeCommand(parts []string) {
	found := false
	for _, c := range h.RootCmd.Commands() {
		if c.Name() == parts[0] || c.HasAlias(parts[0]) {
			// 重新构建完整的命令行参数，包括子命令名称
			// 这样 cobra 才能正确识别和执行子命令
			fullArgs := parts // parts[0] 是命令名，parts[1:] 是参数

			// 重置 rootCmd 的参数并执行
			h.RootCmd.SetArgs(fullArgs)
			err := h.RootCmd.Execute()

			if err != nil {
				// 交互模式：先打印 Usage，再打印错误
				fmt.Println(c.UsageString())
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Unknown command: %s\n", parts[0])
		fmt.Println("Type 'help' to see available commands")
	}
}
