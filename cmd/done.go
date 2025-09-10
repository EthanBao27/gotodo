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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if err := storage.SetDone(id, true); err != nil {
			return err
		}
		fmt.Printf("Task %d marked as done.\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
