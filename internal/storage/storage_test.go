package storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestStorageOperations(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gotodo-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set temporary file path
	testFile := filepath.Join(tempDir, "test_tasks.json")
	SetPath(testFile)

	// Test 1: Add task
	t.Run("AddTask", func(t *testing.T) {
		content := "Test task"
		task, err := Add(content)
		if err != nil {
			t.Errorf("Failed to add task: %v", err)
		}

		if task.ID != 1 {
			t.Errorf("Expected ID 1, got %d", task.ID)
		}

		if task.Content != content {
			t.Errorf("Expected content %s, got %s", content, task.Content)
		}

		if task.Done {
			t.Error("New task should not be marked as done")
		}

		if task.CreatedAt == "" {
			t.Error("Task should have creation time")
		}
	})

	// Test 2: List tasks
	t.Run("ListTasks", func(t *testing.T) {
		tasks, err := List()
		if err != nil {
			t.Errorf("Failed to list tasks: %v", err)
		}

		if len(tasks) != 1 {
			t.Errorf("Expected 1 task, got %d", len(tasks))
		}

		if tasks[0].Content != "Test task" {
			t.Errorf("Expected task content 'Test task', got '%s'", tasks[0].Content)
		}
	})

	// Test 3: Add multiple tasks
	t.Run("AddMultipleTasks", func(t *testing.T) {
		// Add second task
		_, err := Add("Second task")
		if err != nil {
			t.Errorf("Failed to add second task: %v", err)
		}

		// Add third task
		_, err = Add("Third task")
		if err != nil {
			t.Errorf("Failed to add third task: %v", err)
		}

		tasks, err := List()
		if err != nil {
			t.Errorf("Failed to list tasks: %v", err)
		}

		if len(tasks) != 3 {
			t.Errorf("Expected 3 tasks, got %d", len(tasks))
		}

		// Check ID assignment
		if tasks[0].ID != 1 || tasks[1].ID != 2 || tasks[2].ID != 3 {
			t.Error("IDs should be assigned sequentially")
		}
	})

	// Test 4: Mark task as done
	t.Run("MarkTaskDone", func(t *testing.T) {
		err := SetDone(1, true)
		if err != nil {
			t.Errorf("Failed to mark task as done: %v", err)
		}

		tasks, err := List()
		if err != nil {
			t.Errorf("Failed to list tasks: %v", err)
		}

		if !tasks[0].Done {
			t.Error("Task should be marked as done")
		}
	})

	// Test 5: Mark task as undone
	t.Run("MarkTaskUndone", func(t *testing.T) {
		err := SetDone(1, false)
		if err != nil {
			t.Errorf("Failed to mark task as undone: %v", err)
		}

		tasks, err := List()
		if err != nil {
			t.Errorf("Failed to list tasks: %v", err)
		}

		if tasks[0].Done {
			t.Error("Task should be marked as undone")
		}
	})

	// Test 6: Delete task
	t.Run("DeleteTask", func(t *testing.T) {
		err := Delete(2)
		if err != nil {
			t.Errorf("Failed to delete task: %v", err)
		}

		tasks, err := List()
		if err != nil {
			t.Errorf("Failed to list tasks: %v", err)
		}

		if len(tasks) != 2 {
			t.Errorf("Expected 2 tasks after deletion, got %d", len(tasks))
		}

		// Verify correct task was deleted
		for _, task := range tasks {
			if task.ID == 2 {
				t.Error("Task with ID 2 should have been deleted")
			}
		}
	})

	// Test 7: Clear all tasks
	t.Run("ClearAllTasks", func(t *testing.T) {
		err := Clear()
		if err != nil {
			t.Errorf("Failed to clear tasks: %v", err)
		}

		tasks, err := List()
		if err != nil {
			t.Errorf("Failed to list tasks: %v", err)
		}

		if len(tasks) != 0 {
			t.Errorf("Expected 0 tasks after clearing, got %d", len(tasks))
		}
	})
}

func TestErrorCases(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gotodo-test-error")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set temporary file path
	testFile := filepath.Join(tempDir, "test_tasks_error.json")
	SetPath(testFile)

	// Test 1: SetDone with non-existent task
	t.Run("SetDoneNonExistent", func(t *testing.T) {
		err := SetDone(999, true)
		if err == nil {
			t.Error("Expected error when marking non-existent task as done")
		}
	})

	// Test 2: Delete non-existent task
	t.Run("DeleteNonExistent", func(t *testing.T) {
		err := Delete(999)
		if err == nil {
			t.Error("Expected error when deleting non-existent task")
		}
	})

	// Test 3: List with non-existent file (should return empty list)
	t.Run("ListNonExistentFile", func(t *testing.T) {
		// Set path to non-existent file
		nonExistentFile := filepath.Join(tempDir, "non_existent.json")
		SetPath(nonExistentFile)

		tasks, err := List()
		if err != nil {
			t.Errorf("Unexpected error when listing from non-existent file: %v", err)
		}

		if len(tasks) != 0 {
			t.Errorf("Expected empty list from non-existent file, got %d tasks", len(tasks))
		}
	})
}

func TestTaskCreationTime(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gotodo-test-time")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set temporary file path
	testFile := filepath.Join(tempDir, "test_tasks_time.json")
	SetPath(testFile)

	// Record time before adding task
	before := time.Now()

	// Add task
	task, err := Add("Time test task")
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	// Record time after adding task
	after := time.Now()

	// Parse the creation time from task
	createdAt, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", task.CreatedAt)
	if err != nil {
		t.Fatalf("Failed to parse creation time: %v", err)
	}

	// Check that creation time is between before and after
	if createdAt.Before(before) {
		t.Error("Creation time is before task was created")
	}

	if createdAt.After(after) {
		t.Error("Creation time is after task was created")
	}
}

func TestFilePermissions(t *testing.T) {
	// Test with read-only directory
	tempDir, err := os.MkdirTemp("", "gotodo-test-perm")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Make directory read-only
	err = os.Chmod(tempDir, 0444)
	if err != nil {
		t.Fatalf("Failed to change directory permissions: %v", err)
	}

	// Try to use read-only directory
	readOnlyFile := filepath.Join(tempDir, "readonly.json")
	SetPath(readOnlyFile)

	// This should fail due to permission issues
	_, err = Add("Test task")
	if err == nil {
		t.Error("Expected error when writing to read-only directory")
	}

	// Restore permissions for cleanup
	os.Chmod(tempDir, 0755)
}