# Shell Auto-Completion Setup

This guide explains how to set up shell auto-completion for the gotodo CLI application.

## Supported Shells

- **Bash** - Most common Unix/Linux shell
- **Zsh** - Advanced shell with powerful completion
- **Fish** - User-friendly shell with smart completions

## Installation Methods

### Method 1: Quick Setup (Recommended)

#### For Bash
```bash
# Load completion for current session
source <(gotodo completion bash)

# Add to ~/.bashrc for permanent setup
echo 'source <(gotodo completion bash)' >> ~/.bashrc
```

#### For Zsh
```bash
# Load completion for current session
source <(gotodo completion zsh)

# Add to ~/.zshrc for permanent setup
echo 'source <(gotodo completion zsh)' >> ~/.zshrc
```

#### For Fish
```bash
# Load completion for current session
gotodo completion fish | source

# Add to ~/.config/fish/completions/gotodo.fish
gotodo completion fish > ~/.config/fish/completions/gotodo.fish
```

### Method 2: File-based Installation

#### For Bash
```bash
# Generate completion file
gotodo completion bash > /etc/bash_completion.d/gotodo

# Or for user-specific installation
gotodo completion bash > ~/.bash_completion.d/gotodo

# Reload bash
exec bash
```

#### For Zsh
```bash
# Generate completion file in system completions directory
gotodo completion zsh > /usr/local/share/zsh/site-functions/_gotodo

# Or for user-specific installation
gotodo completion zsh > ~/.zsh/completions/_gotodo

# Add to ~/.zshrc if not already there
echo 'fpath=(~/.zsh/completions $fpath)' >> ~/.zshrc
echo 'autoload -U compinit && compinit' >> ~/.zshrc

# Reload zsh
exec zsh
```

#### For Fish
```bash
# Generate completion file
gotodo completion fish > ~/.config/fish/completions/gotodo.fish

# Reload fish or start new session
exec fish
```

## Usage

After installation, auto-completion works automatically:

1. **Command completion**: Type `gotodo` and press `TAB` to see available commands
   ```bash
   gotodo[TAB]
   # Shows: add clear completion config delete done help list
   ```

2. **Flag completion**: Type `-` or `--` and press `TAB` to see available flags
   ```bash
   gotodo list --[TAB]
   # Shows: --done --help --undone
   ```

3. **Context-aware completion**: Completion adapts based on context
   ```bash
   gotodo config [TAB]
   # Shows: set-db show
   ```

## Verification

Test if completion is working:

```bash
# Test command completion
gotodo [TAB][TAB]

# Test flag completion  
gotodo list --[TAB][TAB]

# Test subcommand completion
gotodo config [TAB][TAB]
```

## Troubleshooting

### Bash completion not working

1. Ensure bash-completion is installed:
   ```bash
   # Ubuntu/Debian
   sudo apt-get install bash-completion
   
   # macOS
   brew install bash-completion
   ```

2. Check if completion file is loaded:
   ```bash
   grep -q "gotodo" ~/.bashrc || echo "Completion not loaded"
   ```

3. Reload your shell:
   ```bash
   source ~/.bashrc
   ```

### Zsh completion not working

1. Ensure compinit is loaded in ~/.zshrc:
   ```bash
   echo 'autoload -U compinit && compinit' >> ~/.zshrc
   ```

2. Check completion directory:
   ```bash
   echo $fpath | grep completions
   ```

3. Rebuild completion cache:
   ```bash
   rm -f ~/.zcompdump*
   compinit
   ```

### Fish completion not working

1. Ensure completion directory exists:
   ```bash
   mkdir -p ~/.config/fish/completions
   ```

2. Check if completion file exists:
   ```bash
   ls -la ~/.config/fish/completions/gotodo.fish
   ```

3. Restart fish shell

## System-wide Installation

For system administrators who want to install completion for all users:

#### For Bash
```bash
sudo gotodo completion bash > /etc/bash_completion.d/gotodo
```

#### For Zsh
```bash
sudo gotodo completion zsh > /usr/local/share/zsh/site-functions/_gotodo
```

#### For Fish
```bash
sudo gotodo completion fish > /usr/share/fish/completions/gotodo.fish
```

## Advanced Usage

### Custom completion locations

You can specify custom completion file locations:

```bash
# Custom bash completion
gotodo completion bash > /path/to/custom/completion

# Load it manually
source /path/to/custom/completion
```

### Testing completion without installation

Test completion in current session only:

```bash
# Bash
source <(gotodo completion bash)

# Zsh  
source <(gotodo completion zsh)

# Fish
gotodo completion fish | source
```

## Completion Features

The gotodo completion provides:

- **Command completion**: All available commands
- **Flag completion**: Context-aware flag suggestions
- **Subcommand completion**: Smart subcommand detection
- **File completion**: Where appropriate (e.g., --db flag)
- **Description hints**: Helpful descriptions for options
- **Error handling**: Graceful fallback to file completion

## Integration with Development

For developers working on gotodo:

```bash
# Regenerate completion scripts after changes
./gotodo completion bash > completions/bash
./gotodo completion zsh > completions/zsh
./gotodo completion fish > completions/fish

# Test completion locally
source completions/bash
```

## Tips

1. **Use `double TAB`** to see all available options
2. **Combine with `--help`** for detailed command information
3. **Use arrow keys** to navigate through completion options
4. **Press `Enter`** to select a completion
5. **Press `ESC`** to cancel completion

## Contributing

If you find issues with the completion scripts or want to improve them:

1. Check the Cobra documentation for completion options
2. Test on different shell versions
3. Consider edge cases and error scenarios
4. Update documentation accordingly