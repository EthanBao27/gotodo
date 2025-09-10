/*
Copyright © 2025 Ethan Bao 522425561@qq.com
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/ethanbao27/gotodo/internal/storage"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if err := storage.Delete(id); err != nil {
			return err
		}
		fmt.Printf("Task %d deleted.\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
