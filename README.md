# gotodo

A tiny, delicate todo-cli written in Go.

## Installation

### Install from source

```bash
go install github.com/ethanbao27/gotodo@latest
```

### Build from source

```bash
git clone https://github.com/ethanbao27/gotodo.git
cd gotodo
go build
```

## Usage

### First Time Setup

>[!notice]
> When you run `gotodo` for the first time, you should configure shell completion:

```bash
gotodo init
```

Output:
```
© Copy configuring gotodo completion
Added zsh completion to /Users/username/.zshrc, run 'source /Users/username/.zshrc' to enable.
```

### Basic Commands

#### Add a task

```bash
gotodo add "Task content"
```

#### List all tasks

```bash
gotodo list
```

#### List only done tasks

```bash
gotodo list --done
```

#### List only undone tasks

```bash
gotodo list --undone
```

#### Mark a task as done

```bash
gotodo done <task-id>
```

#### Delete a task

```bash
gotodo delete <task-id>
```

#### Clear all tasks

```bash
gotodo clear --yes
```

#### Use custom database path

```bash
gotodo --db /path/to/custom/file.json <command>
```

## Features

- **Simple and intuitive CLI interface**
- **Automatic shell completion setup** (supports bash, zsh, fish)
- **Task management** (add, list, mark as done, delete)
- **Progress tracking** with visual progress bar
- **Color-coded output** for better readability
- **Persistent storage** using JSON files
- **Cross-platform support** (Linux, macOS, Windows)

## Development

### Requirements

- Go 1.24 or later

### Building

```bash
go build
```

### Testing

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Run tests with race detection
go test -race -v ./...
```

### Linting

```bash
golangci-lint run
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Author

Created by [Ethan Bao](https://github.com/ethanbao27)

<a href="https://github.com/ethanbao27/gotodo/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=ethanbao27/gotodo" />
</a>
