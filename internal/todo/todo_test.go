package todo

import (
	"testing"
	"time"
)

// TestNewTodo verifies that NewTodo creates a todo with correct initial values.
func TestNewTodo(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		title       string
		description string
		wantID      int
		wantTitle   string
		wantDesc    string
		wantErr     bool
	}{
		{
			name:        "create simple todo",
			id:          1,
			title:       "Buy groceries",
			description: "Milk, bread, eggs",
			wantID:      1,
			wantTitle:   "Buy groceries",
			wantDesc:    "Milk, bread, eggs",
			wantErr:     false,
		},
		{
			name:        "create todo with empty description",
			id:          2,
			title:       "Call dentist",
			description: "",
			wantID:      2,
			wantTitle:   "Call dentist",
			wantDesc:    "",
			wantErr:     false,
		},
		{
			name:        "trim whitespace from title",
			id:          3,
			title:       "  Todo with spaces  ",
			description: "  desc with spaces  ",
			wantID:      3,
			wantTitle:   "Todo with spaces",
			wantDesc:    "desc with spaces",
			wantErr:     false,
		},
		{
			name:    "empty title returns error",
			id:      4,
			title:   "",
			wantErr: true,
		},
		{
			name:    "whitespace-only title returns error",
			id:      5,
			title:   "   ",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todo, err := NewTodo(tt.id, tt.title, tt.description)
			if tt.wantErr {
				if err == nil {
					t.Error("NewTodo() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("NewTodo() unexpected error: %v", err)
				return
			}
			if todo.ID != tt.wantID {
				t.Errorf("NewTodo().ID = %v, want %v", todo.ID, tt.wantID)
			}
			if todo.Title != tt.wantTitle {
				t.Errorf("NewTodo().Title = %v, want %v", todo.Title, tt.wantTitle)
			}
			if todo.Description != tt.wantDesc {
				t.Errorf("NewTodo().Description = %v, want %v", todo.Description, tt.wantDesc)
			}
			if todo.Completed {
				t.Errorf("NewTodo().Completed = true, want false")
			}
			if todo.CreatedAt.IsZero() {
				t.Error("NewTodo().CreatedAt is zero")
			}
			if todo.UpdatedAt.IsZero() {
				t.Error("NewTodo().UpdatedAt is zero")
			}
		})
	}
}

// TestTodo_MarkComplete verifies that MarkComplete sets the completed flag and updates the timestamp.
func TestTodo_MarkComplete(t *testing.T) {
	todo, err := NewTodo(1, "Test todo", "")
	if err != nil {
		t.Fatalf("NewTodo() error: %v", err)
	}
	originalTime := todo.UpdatedAt

	// Wait a small amount to ensure time difference
	time.Sleep(time.Millisecond)

	todo.MarkComplete()

	if !todo.Completed {
		t.Error("MarkComplete() did not set Completed to true")
	}
	if !todo.UpdatedAt.After(originalTime) {
		t.Error("MarkComplete() did not update UpdatedAt")
	}
}

// TestTodo_MarkIncomplete verifies that MarkIncomplete clears the completed flag and updates the timestamp.
func TestTodo_MarkIncomplete(t *testing.T) {
	todo, err := NewTodo(1, "Test todo", "")
	if err != nil {
		t.Fatalf("NewTodo() error: %v", err)
	}
	todo.MarkComplete()
	originalTime := todo.UpdatedAt

	// Wait a small amount to ensure time difference
	time.Sleep(time.Millisecond)

	todo.MarkIncomplete()

	if todo.Completed {
		t.Error("MarkIncomplete() did not set Completed to false")
	}
	if !todo.UpdatedAt.After(originalTime) {
		t.Error("MarkIncomplete() did not update UpdatedAt")
	}
}

// TestTodo_Update verifies that Update correctly modifies title and description fields.
func TestTodo_Update(t *testing.T) {
	tests := []struct {
		name          string
		initialTitle  string
		initialDesc   string
		updateTitle   string
		updateDesc    string
		expectedTitle string
		expectedDesc  string
	}{
		{
			name:          "update both title and description",
			initialTitle:  "Old title",
			initialDesc:   "Old description",
			updateTitle:   "New title",
			updateDesc:    "New description",
			expectedTitle: "New title",
			expectedDesc:  "New description",
		},
		{
			name:          "update only title",
			initialTitle:  "Old title",
			initialDesc:   "Keep this",
			updateTitle:   "New title",
			updateDesc:    "",
			expectedTitle: "New title",
			expectedDesc:  "Keep this",
		},
		{
			name:          "update only description",
			initialTitle:  "Keep this",
			initialDesc:   "Old description",
			updateTitle:   "",
			updateDesc:    "New description",
			expectedTitle: "Keep this",
			expectedDesc:  "New description",
		},
		{
			name:          "update with empty strings keeps original",
			initialTitle:  "Original",
			initialDesc:   "Original desc",
			updateTitle:   "",
			updateDesc:    "",
			expectedTitle: "Original",
			expectedDesc:  "Original desc",
		},
		{
			name:          "trim whitespace from updates",
			initialTitle:  "Original",
			initialDesc:   "Original desc",
			updateTitle:   "  New title  ",
			updateDesc:    "  New desc  ",
			expectedTitle: "New title",
			expectedDesc:  "New desc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todo, err := NewTodo(1, tt.initialTitle, tt.initialDesc)
			if err != nil {
				t.Fatalf("NewTodo() error: %v", err)
			}
			originalTime := todo.UpdatedAt

			// Wait a small amount to ensure time difference
			time.Sleep(time.Millisecond)

			todo.Update(tt.updateTitle, tt.updateDesc)

			if todo.Title != tt.expectedTitle {
				t.Errorf("Update() title = %v, want %v", todo.Title, tt.expectedTitle)
			}
			if todo.Description != tt.expectedDesc {
				t.Errorf("Update() description = %v, want %v", todo.Description, tt.expectedDesc)
			}
			if !todo.UpdatedAt.After(originalTime) {
				t.Error("Update() did not update UpdatedAt")
			}
		})
	}
}

// TestTodo_Validate verifies that Validate correctly checks todo data integrity.
func TestTodo_Validate(t *testing.T) {
	tests := []struct {
		name    string
		todo    *Todo
		wantErr bool
	}{
		{
			name: "valid todo",
			todo: &Todo{
				ID:    1,
				Title: "Valid title",
			},
			wantErr: false,
		},
		{
			name: "empty title",
			todo: &Todo{
				ID:    2,
				Title: "",
			},
			wantErr: true,
		},
		{
			name: "whitespace-only title",
			todo: &Todo{
				ID:    3,
				Title: "   ",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.todo.Validate()
			if tt.wantErr && err == nil {
				t.Error("Validate() expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Validate() unexpected error: %v", err)
			}
		})
	}
}
