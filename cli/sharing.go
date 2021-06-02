package cli

import (
	"github.com/spf13/cobra"

	"fmt"
	"log"
)

// Root for sharing/collaboration subcommands
func NewCmdSharingRoot() *cobra.Command {
	gitCmd := &cobra.Command{
		Use:   "sharing",
		Short: "Collaboration commands",
	}

	gitCmd.AddCommand(NewCmdSharingAdd())
	gitCmd.AddCommand(NewCmdSharingRemove())

	return gitCmd
}

// NewCmdSharingAdd generates the Cobra command for "sharing:add"
func NewCmdSharingAdd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add EMAIL",
		Short: "Add a collaborator",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("Please specify at least one collaborator")
			}

			cc := cmd.Context()
			c, err := newAPIClient(cc)
			if err != nil {
				return err
			}

			for _, name := range args {
				err := c.AddCollaborator(cc, name)

				if err != nil {
					log.Printf("Problem adding %q: %s\n", name, err)
					continue
				}

				fmt.Printf("Invited %q as a collaborator\n", name)
			}

			return nil
		},
	}

	return addCmd
}

// NewCmdSharingRemove generates the Cobra command for "sharing:add"
func NewCmdSharingRemove() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "remove EMAIL",
		Short: "Remove a collaborator",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("Please specify at least one collaborator")
			}

			cc := cmd.Context()
			c, err := newAPIClient(cc)
			if err != nil {
				return err
			}

			for _, name := range args {
				err := c.RemoveCollaborator(cc, name)

				if err != nil {
					log.Printf("Problem removing %q: %s\n", name, err)
					continue
				}

				fmt.Printf("Removed %q as a collaborator\n", name)
			}

			return nil
		},
	}

	return addCmd
}
