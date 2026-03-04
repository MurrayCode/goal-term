package suggest

import (
	"strings"
	"testing"

	"github.com/murraycode/goal-term/internal/goal"
	"github.com/murraycode/goal-term/internal/storage"
)

func TestBuildPromptIncludesGoalAndTasks(t *testing.T) {
	cfg := storage.Config{
		Goal: "Ship v1",
		Tasks: []goal.Task{
			{Title: "Plan", Status: goal.StatusTodo},
			{Title: "Execute", Status: goal.StatusDone},
		},
	}

	prompt := buildPrompt(cfg)
	if !strings.Contains(prompt, "Goal: Ship v1") {
		t.Fatalf("expected goal in prompt")
	}
	if !strings.Contains(prompt, "1. Plan (todo)") {
		t.Fatalf("expected first task in prompt")
	}
	if !strings.Contains(prompt, "2. Execute (done)") {
		t.Fatalf("expected second task in prompt")
	}
}
