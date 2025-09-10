package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/ethanbao27/gotodo/internal/storage"
	"github.com/ethanbao27/gotodo/internal/ui"
	"github.com/fatih/color"
)

// connect to friend and fetch tasks
func ConnectAndFetch(addr string) error {
	port := ":8088"
	conn, err := net.Dial("tcp", addr+port)
	if err != nil {
		return fmt.Errorf("dial error: %w", err)
	}
	defer conn.Close()

	// send request
	fmt.Fprintf(conn, "GET_TODOS\n")

	// read response
	resp, err := io.ReadAll(conn)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}

	var tasks []storage.Task
	if err := json.Unmarshal(resp, &tasks); err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	if len(tasks) == 0 {
		color.New(color.FgYellow).Println("No tasks received from friend.")
		return nil
	}

	// ==== Header ====
	fmt.Println()
	color.New(color.FgBlue, color.Bold).Printf(" Friend's Todo List @ %s\n", addr)
	color.New(color.FgWhite, color.Faint).Printf("  %d tasks\n\n", len(tasks))

	// ==== Stats ====
	doneCount := 0
	totalCount := len(tasks)
	for _, t := range tasks {
		if t.Done {
			doneCount++
		}
	}
	progress := float64(doneCount) / float64(totalCount) * 100

	// ==== Progress bar ====
	ui.PrintProgressBar(progress)
	fmt.Println()

	// ==== Tasks ====
	for _, t := range tasks {
		var statusIcon string
		var statusColor color.Attribute
		if t.Done {
			statusIcon = "[✓]"
			statusColor = color.FgGreen
		} else {
			statusIcon = "[ ]"
			statusColor = color.FgWhite
		}

		color.New(statusColor).Printf(" %s %3d ", statusIcon, t.ID)
		color.New(color.FgWhite).Print(t.Content)
		if t.CreatedAt != "" {
			color.New(color.FgCyan, color.Faint).Printf("  (%s)", t.CreatedAt)
		}
		fmt.Println()
	}

	fmt.Println()
	ui.PrintProgressSummary(doneCount, totalCount, progress)
	return nil
}
