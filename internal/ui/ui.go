package ui

import (
	"fmt"
	"io"

	"github.com/murraycode/goal-term/internal/storage"
)

func PrintGoal(w io.Writer, cfg storage.Config) {
	fmt.Fprintf(w, "Goal: %s\n", cfg.Goal)
	for i, task := range cfg.Tasks {
		fmt.Fprintf(w, "%d. %s (%s)\n", i+1, task.Title, task.Status)
	}
}

func PrintSuggestions(w io.Writer, suggestions string) {
	if suggestions == "" {
		return
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Suggestions:")
	fmt.Fprintln(w, suggestions)
}
