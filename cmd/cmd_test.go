package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ethanbao27/gotodo/internal/storage"
)

func TestAddCommand(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gotodo-cmd-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set temporary file path
	testFile := filepath.Join(tempDir, "test_tasks.json")
	storage.SetPath(testFile)

	// Test adding a single task
	t.Run("AddSingleTask", func(t *testing.T) {
		// Execute add command with database path
		// Create a temporary root command for testing
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "add", "Test task"})
		err := testRootCmd.Execute()
		if err != nil {
			t.Errorf("Add command failed: %v", err)
		}

		// Set the storage path to test file for verification
		storage.SetPath(testFile)
		
		// Verify task was added
		tasks, err := storage.List()
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

	// Test adding task with no arguments (should fail)
	t.Run("AddNoArgs", func(t *testing.T) {
		// Execute add command with no arguments
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"add"})
		err := testRootCmd.Execute()
		if err == nil {
			t.Error("Expected error when adding task with no arguments")
		}
	})
}

func TestListCommand(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gotodo-cmd-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set temporary file path
	testFile := filepath.Join(tempDir, "test_tasks.json")
	storage.SetPath(testFile)

	// Test listing empty tasks
	t.Run("ListEmpty", func(t *testing.T) {
		// Execute list command
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "list"})
		err := testRootCmd.Execute()
		if err != nil {
			t.Errorf("List command failed: %v", err)
		}
	})

	// Test listing with tasks
	t.Run("ListWithTasks", func(t *testing.T) {
		// Add some tasks first
		_, err := storage.Add("First task")
		if err != nil {
			t.Fatalf("Failed to add test task: %v", err)
		}
		_, err = storage.Add("Second task")
		if err != nil {
			t.Fatalf("Failed to add test task: %v", err)
		}

		// Execute list command
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "list"})
		err = testRootCmd.Execute()
		if err != nil {
			t.Errorf("List command failed: %v", err)
		}
	})
}

func TestDoneCommand(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gotodo-cmd-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set temporary file path
	testFile := filepath.Join(tempDir, "test_tasks.json")
	storage.SetPath(testFile)

	// Test marking task as done
	t.Run("MarkTaskDone", func(t *testing.T) {
		_, err := storage.Add("Task to complete")
		if err != nil {
			t.Fatalf("Failed to add test task: %v", err)
		}

		// Execute done command
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "done", "1"})
		err = testRootCmd.Execute()
		if err != nil {
			t.Errorf("Done command failed: %v", err)
		}

		// Verify task was marked as done
		tasks, err := storage.List()
		if err != nil {
			t.Errorf("Failed to list tasks: %v", err)
		}

		if !tasks[0].Done {
			t.Error("Task should be marked as done")
		}
	})

	// Test marking non-existent task as done
	t.Run("MarkNonExistentDone", func(t *testing.T) {
		// Execute done command with non-existent ID
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "done", "999"})
		err := testRootCmd.Execute()
		if err == nil {
			t.Error("Expected error when marking non-existent task as done")
		}
	})

	// Test marking task as done with invalid ID
	t.Run("MarkInvalidID", func(t *testing.T) {
		// Execute done command with invalid ID
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "done", "invalid"})
		err := testRootCmd.Execute()
		if err == nil {
			t.Error("Expected error when marking task with invalid ID as done")
		}
	})
}

func TestDeleteCommand(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gotodo-cmd-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set temporary file path
	testFile := filepath.Join(tempDir, "test_tasks.json")
	storage.SetPath(testFile)

	// Test deleting a task
	t.Run("DeleteTask", func(t *testing.T) {
		_, err := storage.Add("Task to delete")
		if err != nil {
			t.Fatalf("Failed to add test task: %v", err)
		}

		// Execute delete command
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "delete", "1"})
		err = testRootCmd.Execute()
		if err != nil {
			t.Errorf("Delete command failed: %v", err)
		}

		// Verify task was deleted
		tasks, err := storage.List()
		if err != nil {
			t.Errorf("Failed to list tasks: %v", err)
		}

		if len(tasks) != 0 {
			t.Error("Task should have been deleted")
		}
	})

	// Test deleting non-existent task
	t.Run("DeleteNonExistent", func(t *testing.T) {
		// Execute delete command with non-existent ID
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "delete", "999"})
		err := testRootCmd.Execute()
		if err == nil {
			t.Error("Expected error when deleting non-existent task")
		}
	})
}

func TestClearCommand(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gotodo-cmd-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set temporary file path
	testFile := filepath.Join(tempDir, "test_tasks.json")
	storage.SetPath(testFile)

	// Test clearing all tasks without confirmation (should fail)
	t.Run("ClearWithoutConfirmation", func(t *testing.T) {
		_, err := storage.Add("Task to clear")
		if err != nil {
			t.Fatalf("Failed to add test task: %v", err)
		}

		// Execute clear command without --yes flag
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "clear"})
		err = testRootCmd.Execute()
		if err == nil {
			t.Error("Expected error when clearing without confirmation")
		}
	})

	// Test clearing all tasks with confirmation
	t.Run("ClearWithConfirmation", func(t *testing.T) {
		_, err := storage.Add("Task to clear")
		if err != nil {
			t.Fatalf("Failed to add test task: %v", err)
		}

		// Execute clear command with --yes flag
		testRootCmd := rootCmd
		testRootCmd.SetArgs([]string{"--db", testFile, "clear", "--yes"})
		err = testRootCmd.Execute()
		if err != nil {
			t.Errorf("Clear command failed: %v", err)
		}

		// Verify all tasks were cleared
		tasks, err := storage.List()
		if err != nil {
			t.Errorf("Failed to list tasks: %v", err)
		}

		if len(tasks) != 0 {
			t.Error("All tasks should have been cleared")
		}
	})
}