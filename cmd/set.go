package cmd

import (
	"github.com/spf13/cobra"
)

var setKeyValueCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value for a key",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
