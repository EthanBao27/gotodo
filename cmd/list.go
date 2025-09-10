/*
Copyright © 2025 Ethan Bao 522425561@qq.com
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethanbao27/gotodo/internal/storage"
	"github.com/ethanbao27/gotodo/internal/ui"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var onlyDone bool
var onlyUndone bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.List()
		if err != nil {
			return err
		}
		if len(tasks) == 0 {
			color.New(color.FgYellow).Println("No tasks.")
			return nil
		}

		// Calculate statistics
		doneCount := 0
		totalCount := len(tasks)
		for _, t := range tasks {
			if t.Done {
				doneCount++
			}
		}
		progress := float64(doneCount) / float64(totalCount) * 100

		// Print minimal header
		fmt.Println()
		color.New(color.FgBlue, color.Bold).Printf("  TASKS  ")
		color.New(color.FgWhite, color.Faint).Printf("  %d total, %d done\n", totalCount, doneCount)
		fmt.Println()

		// Print progress bar with animation
		ui.PrintProgressBar(progress)

		fmt.Println()

		// Print tasks
		for _, t := range tasks {
			if onlyDone && !t.Done {
				continue
			}
			if onlyUndone && t.Done {
				continue
			}

			// Parse and format creation time
			createdAt := "Unknown"
			if t.CreatedAt != "" {
				if parsedTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", t.CreatedAt); err == nil {
					createdAt = parsedTime.Format("Jan 02 15:04")
				} else if parsedTime, err := time.Parse("2006-01-02 15:04:05 -0700 MST", t.CreatedAt); err == nil {
					createdAt = parsedTime.Format("Jan 02 15:04")
				} else {
					createdAt = strings.Split(t.CreatedAt, ".")[0]
					if len(createdAt) > 16 {
						createdAt = createdAt[:16]
					}
				}
			}

			// Minimal format: [✓] ID Content (date)
			var statusIcon string
			var statusColor color.Attribute
			if t.Done {
				statusIcon = "[✓]"
				statusColor = color.FgGreen
			} else {
				statusIcon = "[ ]"
				statusColor = color.FgWhite
			}

			// Print task with minimal styling
			color.New(statusColor).Printf(" %s %3d ", statusIcon, t.ID)
			color.New(color.FgWhite).Print(t.Content)
			color.New(color.FgCyan, color.Faint).Printf("  %s\n", createdAt)
		}

		fmt.Println()
		ui.PrintProgressSummary(doneCount, totalCount, progress)
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVar(&onlyDone, "done", false, "show done only")
	listCmd.Flags().BoolVar(&onlyUndone, "undone", false, "show undone only")
	rootCmd.AddCommand(listCmd)
}
