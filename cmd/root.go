package cmd

import (
	"context"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{}

func init() {
	root.AddCommand(setKeyValueCmd)
}

// Exec runs the command registered with root
func Exec(input string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	args := strings.Split(input, " ")

	root.SetArgs(args)
	return root.ExecuteContext(ctx)
}
