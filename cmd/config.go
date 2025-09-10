/*
Copyright © 2025 Ethan Bao 522425561@qq.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure gotodo settings",
	Long:  `Configure gotodo settings like database path.`,
}

var setDbPathCmd = &cobra.Command{
	Use:   "set-db",
	Short: "Set permanent database path",
	Long:  `Set a permanent database path that will be used for all future commands.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		
		// Create directory if it doesn't exist
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
		
		// Test if path is writable
		testFile := filepath.Join(dir, ".gotodo_test")
		if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
			return fmt.Errorf("path is not writable: %v", err)
		}
		os.Remove(testFile)
		
		// Save config to ~/.gotodo/config.json
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %v", err)
		}
		
		configDir := filepath.Join(home, ".gotodo")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return fmt.Errorf("failed to create config directory: %v", err)
		}
		
		configFile := filepath.Join(configDir, "config.json")
		config := map[string]string{"db_path": path}
		
		// Save config
		data, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			return fmt.Errorf("failed to marshal config: %v", err)
		}
		
		if err := os.WriteFile(configFile, data, 0644); err != nil {
			return fmt.Errorf("failed to save config: %v", err)
		}
		
		color.New(color.FgGreen).Printf("✓ Database path set to: %s\n", path)
		color.New(color.FgYellow).Printf("Note: Restart gotodo to take effect\n")
		return nil
	},
}

var showConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %v", err)
		}
		
		configFile := filepath.Join(home, ".gotodo", "config.json")
		
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			color.New(color.FgYellow).Println("No configuration file found, using default settings")
			return nil
		}
		
		data, err := os.ReadFile(configFile)
		if err != nil {
			return fmt.Errorf("failed to read config file: %v", err)
		}
		
		var config map[string]string
		if err := json.Unmarshal(data, &config); err != nil {
			return fmt.Errorf("failed to parse config file: %v", err)
		}
		
		color.New(color.FgGreen).Println("Current configuration:")
		if dbPath, exists := config["db_path"]; exists {
			color.New(color.FgCyan).Printf("Database path: %s\n", dbPath)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setDbPathCmd)
	configCmd.AddCommand(showConfigCmd)
}