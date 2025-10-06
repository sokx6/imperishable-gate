// ensure_token.go
package service

import (
	"errors"
	"fmt"
	"os"

	"imperishable-gate/internal/client/utils"
)

func EnsureValidTokenWithPrompt(addr, accessToken string) (string, error) {
	token, err := GetTokenAutomatically(addr, accessToken)

	if err != nil {
		switch {
		case errors.Is(err, utils.ErrNoRefreshToken):
			fmt.Fprintln(os.Stdout, "Authentication expired: no refresh token found.")
			fmt.Fprintln(os.Stdout, "Please run 'login' to sign in again.")
		default:
			fmt.Fprintf(os.Stderr, "Failed to obtain valid token: %v\n", err)
			return "", fmt.Errorf("failed to get valid token: %w", err)
		}

		return "", err
	}

	return token, nil
}
