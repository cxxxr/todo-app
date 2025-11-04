// Package todo provides the core domain model for todo items.
//
// This package contains the fundamental Todo data structure with JSON tags
// for serialization via encoding/json. A Storage interface for persistence
// abstraction will be implemented in the next PR (Phase 2: Storage Layer).
package todo

import (
	"errors"
	"strings"
	"time"
)

var (
	// ErrEmptyTitle is returned when a todo title is empty or only whitespace.
	ErrEmptyTitle = errors.New("todo title cannot be empty")
)

// Todo represents a single todo item with its metadata.
// JSON tags enable serialization via encoding/json Marshal/Unmarshal.
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewTodo creates a new Todo with the given title and description.
// Returns ErrEmptyTitle if the title is empty or only whitespace.
func NewTodo(id int, title, description string) (*Todo, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, ErrEmptyTitle
	}

	now := time.Now()
	return &Todo{
		ID:          id,
		Title:       title,
		Description: strings.TrimSpace(description),
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// MarkComplete marks the todo as completed and updates the timestamp.
func (t *Todo) MarkComplete() {
	t.Completed = true
	t.UpdatedAt = time.Now()
}

// MarkIncomplete marks the todo as incomplete and updates the timestamp.
func (t *Todo) MarkIncomplete() {
	t.Completed = false
	t.UpdatedAt = time.Now()
}

// Update updates the todo's title and description, and refreshes the timestamp.
// Empty strings are ignored to preserve existing values.
func (t *Todo) Update(title, description string) {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)

	if title != "" {
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}

	t.UpdatedAt = time.Now()
}

// Validate checks if the todo has valid data.
func (t *Todo) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return ErrEmptyTitle
	}
	return nil
}
