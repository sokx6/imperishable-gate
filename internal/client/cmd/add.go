package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"

	types "imperishable-gate/internal"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new link to the server database",
	RunE: func(cmd *cobra.Command, args []string) error {
		host, _ := cmd.Flags().GetString("host")
		link, _ := cmd.Flags().GetString("link")

		// 构造请求体
		reqBody := types.AddRequest{
			Action: "add",
			Link:   link,
		}

		body, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}

		url := fmt.Sprintf("http://%s", host)
		fmt.Printf("-- Requesting POST method to %s with payload\n", host)
		fmt.Printf("%s\n", body)

		// 发起 POST 请求
		resp, err := http.Post(url, "application/json", strings.NewReader(string(body)))
		if err != nil {
			return fmt.Errorf("request failed: %w", err)
		}

		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}

		var result types.PingResponse
		if err := json.Unmarshal(respBody, &result); err != nil {
			return fmt.Errorf("invalid JSON response: %w", err)
		}

		fmt.Printf("-- Receiving response\n%s\n", respBody)

		return nil
	},
}

func init() {
	addCmd.Flags().StringP("host", "H", "127.0.0.1:8080", "Server host:port to send add request")
	addCmd.Flags().StringP("link", "l", "", "link to add")
	rootCmd.AddCommand(addCmd)
}
