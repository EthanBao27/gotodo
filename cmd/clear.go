/*
Copyright © 2025 Ethan Bao 522425561@qq.com
*/
package cmd

import (
	"fmt"

	"github.com/ethanbao27/gotodo/internal/storage"
	"github.com/spf13/cobra"
)

var yes bool

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !yes {
			return fmt.Errorf("this will remove ALL tasks; confirm with --yes")
		}
		if err := storage.Clear(); err != nil {
			return err
		}
		fmt.Println("All tasks cleared.")
		return nil
	},
}

func init() {
	clearCmd.Flags().BoolVar(&yes, "yes", false, "confirm clearing all tasks")
	rootCmd.AddCommand(clearCmd)
}
