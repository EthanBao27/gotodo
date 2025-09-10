package ui

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

func PrintProgressBar(progress float64) {
	width := 40
	filled := int(math.Round(float64(width) * progress / 100))

	// Animated progress bar
	fmt.Print("  ")
	// color.New(color.FgBlue).Print("[")

	for i := 0; i < width; i++ {
		if i < filled {
			color.New(color.FgGreen).Print("█")
		} else {
			color.New(color.FgWhite, color.Faint).Print("░")
		}
	}

	// color.New(color.FgBlue).Print("]")
	color.New(color.FgWhite).Printf(" %5.1f%%\n", progress)
}

func PrintProgressSummary(done, total int, progress float64) {
	// Minimal summary with subtle animation
	color.New(color.FgWhite, color.Bold).Print("  Status: ")

	if progress == 100 {
		color.New(color.FgGreen).Print("✓ Complete")
	} else if progress >= 75 {
		color.New(color.FgYellow).Print("◐ Nearly done")
	} else if progress >= 50 {
		color.New(color.FgYellow).Print("◑ Halfway")
	} else if progress >= 25 {
		color.New(color.FgBlue).Print("◒ In progress")
	} else {
		color.New(color.FgBlue).Print("◓ Just started")
	}

	color.New(color.FgWhite, color.Faint).Printf("  (%d/%d)\n", done, total)
}
