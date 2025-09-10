/*
Copyright © 2025 Ethan Bao 522425561@qq.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/ethanbao27/gotodo/internal/storage"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var dbPath string

var rootCmd = &cobra.Command{
	Use:   "gotodo",
	Short: "A tiny,delicate todo-cli written in Go",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		usr, _ := user.Current()
		marker := filepath.Join(usr.HomeDir, ".gotodo", "init_done")

		if _, err := os.Stat(marker); os.IsNotExist(err) {
			fmt.Println("© Copy configuring gotodo completion")
			if err := InitSetup(); err != nil {
				return err
			}
			if err := os.MkdirAll(filepath.Dir(marker), 0755); err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
			if err := os.WriteFile(marker, []byte("done"), 0644); err != nil {
				return fmt.Errorf("failed to write marker file: %v", err)
			}
		}
		// Load config if no --db flag is provided
		if dbPath == "" {
			home, err := os.UserHomeDir()
			if err == nil {
				configFile := filepath.Join(home, ".gotodo", "config.json")
				if data, err := os.ReadFile(configFile); err == nil {
					var config map[string]string
					if json.Unmarshal(data, &config) == nil {
						if configuredPath, exists := config["db_path"]; exists {
							dbPath = configuredPath
						}
					}
				}
			}
		}

		if dbPath != "" {
			storage.SetPath(dbPath)
		} else {
			storage.SetPath(storage.GetCurrentPath())
		}

		// Only show database path for certain commands
		// Get the full command path to check parent commands
		fullCmd := cmd.CommandPath()
		shouldShowPath := true

		// Don't show path for list, config, and completion commands
		if fullCmd == "gotodo list" || fullCmd == "gotodo config" ||
			fullCmd == "gotodo config show" || fullCmd == "gotodo config set-db" ||
			fullCmd == "gotodo completion" {
			shouldShowPath = false
		}

		if shouldShowPath {
			currentPath := storage.GetCurrentPath()
			color.New(color.FgCyan).Printf("Using database path: %s\n", currentPath)
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbPath, "db", "", "path to store tasks")
}
