package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	RunE: func(cmd *cobra.Command, args []string) error {
		return InitSetup()
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

func init() {
	rootCmd.AddCommand(initCmd)
}
