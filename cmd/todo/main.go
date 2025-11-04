// Package main provides the CLI entry point for the todo application.
//
// Phase 1: Basic command routing. Additional commands (add, list, complete,
// delete) with output formatting will be added in subsequent PRs.
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		// Write error to stderr. If this fails, there's nothing we can do
		// as we're about to exit anyway.
		if _, writeErr := fmt.Fprintf(os.Stderr, "Error: %v\n", err); writeErr != nil {
			// Can't write to stderr, just exit with error code
			os.Exit(1)
		}
		os.Exit(1)
	}
}

// run is the main entry point for the CLI application.
// It handles command parsing and execution, returning any errors
// that occur during operation.
func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("no command specified. Use 'todo help' for usage information")
	}

	command := os.Args[1]
	switch command {
	case "help", "--help", "-h":
		return printHelp()
	default:
		return fmt.Errorf("unknown command: %s. Use 'todo help' for usage information", command)
	}
}

// printHelp displays usage information for the CLI.
func printHelp() error {
	_, err := fmt.Println(`todo - A simple CLI todo application (v0.1.0 - initial setup)

Usage:
  todo <command> [arguments]

Available Commands:
  help        Show this help message

Additional commands (add, list, complete, delete) will be available in future versions.`)
	if err != nil {
		return fmt.Errorf("failed to write help output: %w", err)
	}
	return nil
}
