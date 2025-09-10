# Simple zsh completion for gotodo
_gotodo() {
  local -a commands
  commands=(
    'add:Add a new task'
    'list:List all tasks'
    'done:Mark a task as done'
    'delete:Delete a task'
    'clear:Clear all tasks'
    'config:Configure gotodo settings'
    'completion:Generate shell completion script'
    'help:Help about any command'
    'init:Initialize shell completion'
  )

  if (( CURRENT == 2 )); then
    _describe 'command' commands
  else
    # Handle subcommands and flags
    case $words[2] in
      add)
        _message 'task content'
        ;;
      list)
        _arguments '--done[Show only done tasks]' '--undone[Show only undone tasks]'
        ;;
      done|delete)
        # Show task IDs for completion
        local -a task_ids
        if [[ -f ~/.gotodo/tasks.json ]]; then
          task_ids=($(jq -r '.[].id' ~/.gotodo/tasks.json 2>/dev/null))
        fi
        _describe 'task id' task_ids
        ;;
      config)
        _arguments '1: :(show set-db)'
        ;;
      completion)
        _arguments '1: :(bash zsh fish powershell)'
        ;;
    esac
  fi
}

compdef _gotodo gotodo