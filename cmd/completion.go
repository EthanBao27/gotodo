package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [shell]",
	Short: "Generate shell completion script",
	Long: `To load completions:

Bash:
  $ gotodo completion bash

  # To load completions for each session, execute once:
  # Linux:
  $ gotodo completion bash > /etc/bash_completion.d/gotodo
  # macOS:
  $ gotodo completion bash > /usr/local/etc/bash_completion.d/gotodo

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ gotodo completion zsh > "${fpath[1]}/_gotodo"

  # You will need to start a new shell for this setup to take effect.

fish:
  $ gotodo completion fish | source

  # To load completions for each session, execute once:
  $ gotodo completion fish > ~/.config/fish/completions/gotodo.fish

PowerShell:
  PS> gotodo completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> gotodo completion powershell > gotodo.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		shell := args[0]
		return installCompletion(shell)
	},
}

func installCompletion(shell string) error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}

	switch shell {
	case "bash":
		return installBashCompletion(usr.HomeDir)
	case "zsh":
		return installZshCompletion(usr.HomeDir)
	case "fish":
		return installFishCompletion(usr.HomeDir)
	case "powershell":
		return installPowerShellCompletion(usr.HomeDir)
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}
}

func installBashCompletion(homeDir string) error {
	// Try system-wide installation first
	completionDirs := []string{
		"/etc/bash_completion.d",
		"/usr/local/etc/bash_completion.d",
		filepath.Join(homeDir, ".bash_completion.d"),
	}

	var targetPath string
	for _, dir := range completionDirs {
		if err := os.MkdirAll(dir, 0755); err == nil {
			targetPath = filepath.Join(dir, "gotodo")
			break
		}
	}

	if targetPath == "" {
		return fmt.Errorf("could not find writable bash completion directory")
	}

	file, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("failed to create completion file: %v", err)
	}
	defer file.Close()

	if err := rootCmd.GenBashCompletion(file); err != nil {
		return fmt.Errorf("failed to generate bash completion: %v", err)
	}

	fmt.Printf("Bash completion installed to: %s\n", targetPath)
	fmt.Printf("Run 'source ~/.bashrc' or start a new shell to enable completion.\n")
	return nil
}

func installZshCompletion(homeDir string) error {
	// Try system-wide installation first
	completionDirs := []string{
		"/usr/local/share/zsh/site-functions",
		"/usr/share/zsh/site-functions",
		filepath.Join(homeDir, ".zsh", "completions"),
		filepath.Join(homeDir, ".oh-my-zsh", "completions"),
	}

	var targetPath string
	var lastErr error
	for _, dir := range completionDirs {
		if err := os.MkdirAll(dir, 0755); err == nil {
			// Test if we can write to this directory
			testFile := filepath.Join(dir, ".test_write")
			if err := os.WriteFile(testFile, []byte("test"), 0644); err == nil {
				os.Remove(testFile)
				targetPath = filepath.Join(dir, "_gotodo")
				break
			} else {
				lastErr = err
			}
		} else {
			lastErr = err
		}
	}

	if targetPath == "" {
		return fmt.Errorf("could not find writable zsh completion directory, last error: %v", lastErr)
	}

	file, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("failed to create completion file: %v", err)
	}
	defer file.Close()

	if err := rootCmd.GenZshCompletion(file); err != nil {
		return fmt.Errorf("failed to generate zsh completion: %v", err)
	}

	fmt.Printf("Zsh completion installed to: %s\n", targetPath)
	fmt.Printf("Run 'source ~/.zshrc' or start a new shell to enable completion.\n")
	return nil
}

func installFishCompletion(homeDir string) error {
	completionDir := filepath.Join(homeDir, ".config", "fish", "completions")
	if err := os.MkdirAll(completionDir, 0755); err != nil {
		return fmt.Errorf("failed to create fish completions directory: %v", err)
	}

	targetPath := filepath.Join(completionDir, "gotodo.fish")
	file, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("failed to create completion file: %v", err)
	}
	defer file.Close()

	if err := rootCmd.GenFishCompletion(file, true); err != nil {
		return fmt.Errorf("failed to generate fish completion: %v", err)
	}

	fmt.Printf("Fish completion installed to: %s\n", targetPath)
	fmt.Printf("Start a new fish shell to enable completion.\n")
	return nil
}

func installPowerShellCompletion(homeDir string) error {
	targetPath := filepath.Join(homeDir, "gotodo.ps1")
	file, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("failed to create completion file: %v", err)
	}
	defer file.Close()

	if err := rootCmd.GenPowerShellCompletionWithDesc(file); err != nil {
		return fmt.Errorf("failed to generate powershell completion: %v", err)
	}

	fmt.Printf("PowerShell completion installed to: %s\n", targetPath)
	fmt.Printf("Add the following to your PowerShell profile:\n")
	fmt.Printf("    %s\n", targetPath)
	return nil
}

func init() {
	rootCmd.AddCommand(completionCmd)
}