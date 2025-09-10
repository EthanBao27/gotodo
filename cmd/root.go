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
	"strings"

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

func InitSetup() error {
	usr, _ := user.Current()
	shell := os.Getenv("SHELL")

	switch filepath.Base(shell) {
	case "bash":
		rc := filepath.Join(usr.HomeDir, ".bashrc")
		added, err := appendIfMissing(rc, "\n# gotodo completion\nsource <(gotodo completion bash)\n",
			"gotodo completion bash")
		if err != nil {
			return err
		}
		if added {
			fmt.Printf("Added bash completion to %s, run 'source %s' to enable.\n", rc, rc)
		}
		return nil
	case "zsh":
		rc := filepath.Join(usr.HomeDir, ".zshrc")
		added, err := appendIfMissing(rc, "\n# gotodo completion\nsource <(gotodo completion zsh)\n",
			"gotodo completion zsh")
		if err != nil {
			return err
		}
		if added {
			fmt.Printf("Added zsh completion to %s, run 'source %s' to enable.\n", rc, rc)
		}
		return nil
	case "fish":
		dir := filepath.Join(usr.HomeDir, ".config/fish/completions")
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(dir, "gotodo.fish"))
		if err != nil {
			return err
		}
		defer f.Close()
		err = generateFishCompletion(f)
		if err == nil {
			fmt.Printf("Added fish completion to %s\n", filepath.Join(dir, "gotodo.fish"))
		}
		return err
	default:
		fmt.Printf("Shell %s not supported, please set up manually.\n", shell)
	}
	return nil
}

func generateFishCompletion(f *os.File) error {
	// Create a temporary copy of rootCmd to generate completion
	tempRootCmd := &cobra.Command{
		Use:   "gotodo",
		Short: "A tiny,delicate todo-cli written in Go",
	}
	return tempRootCmd.GenFishCompletion(f, true)
}

func appendIfMissing(rc, snippet, marker string) (bool, error) {
	b, _ := os.ReadFile(rc)
	if strings.Contains(string(b), marker) {
		return false, nil
	}
	f, err := os.OpenFile(rc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()
	_, err = f.WriteString(snippet)
	if err != nil {
		return false, err
	}
	return true, nil
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
