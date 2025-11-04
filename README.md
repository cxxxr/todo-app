# Todo App

A simple command-line todo application written in Go.

## Development Status

This project is being developed incrementally in small PRs:

- ✅ **Phase 1**: Project setup and core domain model (current)
- 🔄 **Phase 2**: Storage layer with JSON persistence (planned)
- 🔄 **Phase 3**: Basic CLI commands (add, list) (planned)
- 🔄 **Phase 4**: Additional CLI commands (complete, delete) (planned)

## Installation

```bash
go build -o todo ./cmd/todo
```

## Usage

```bash
# Show help
./todo help
```

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o todo ./cmd/todo
```
