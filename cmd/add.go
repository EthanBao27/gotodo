/*
Copyright © 2025 Ethan Bao 522425561@qq.com
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/ethanbao27/gotodo/internal/storage"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <task>",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content := strings.Join(args, " ")
		t, err := storage.Add(content)
		if err != nil {
			return err
		}
		fmt.Printf("Added [%d] %s\n", t.ID, t.Content)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
