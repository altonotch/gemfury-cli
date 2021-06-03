package cli

import (
	"github.com/gemfury/cli/api"
	"github.com/spf13/cobra"

	"context"
	"fmt"
)

// NewCmdWhoAmI generates the Cobra command for "whoami"
func NewCmdWhoAmI() *cobra.Command {
	whoCmd := &cobra.Command{
		Use:   "whoami",
		Short: "Return current account",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := whoAMI(cmd.Context())
			if err != nil {
				return err
			}

			fmt.Printf("You are logged in as %q\n", resp.Name)
			return nil
		},
	}

	return whoCmd
}

func whoAMI(cc context.Context) (*api.AccountResponse, error) {
	c, err := newAPIClient(cc)
	if err != nil {
		return nil, err
	}

	return c.WhoAmI(cc)
}
