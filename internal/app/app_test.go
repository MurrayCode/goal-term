package app

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/murraycode/goal-term/internal/cli"
	"github.com/murraycode/goal-term/internal/goal"
	"github.com/murraycode/goal-term/internal/storage"
)

type stubSuggester struct {
	text string
	Err  error
}

func (s stubSuggester) Suggest(_ context.Context, _ storage.Config) (string, error) {
	return s.text, s.Err
}

func TestRunListShowsSuggestions(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cfg := storage.Config{
		Goal:  "Ship v1",
		Tasks: []goal.Task{{Title: "Finish", Status: goal.StatusTodo}},
	}
	path := t.TempDir() + "/goalterm.json"
	if err := storage.SaveConfig(path, cfg); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	err := Run(context.Background(), []string{"list"}, Env{
		ConfigPath: path,
		Out:        &out,
		Err:        &errOut,
		Suggester:  stubSuggester{text: "Try smaller steps"},
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if !bytes.Contains(out.Bytes(), []byte("Suggestions:")) {
		t.Fatalf("expected suggestions in output")
	}
}

func TestRunCompleteOutOfRange(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cfg := storage.Config{Goal: "Ship v1"}
	path := t.TempDir() + "/goalterm.json"
	if err := storage.SaveConfig(path, cfg); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	err := Run(context.Background(), []string{"complete", "1"}, Env{
		ConfigPath: path,
		Out:        &out,
		Err:        &errOut,
	})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, cli.ErrUsage) {
		t.Fatalf("expected usage error, got %v", err)
	}
}

func TestRunSetGoal(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cfg := storage.Config{Goal: "Old"}
	path := t.TempDir() + "/goalterm.json"
	if err := storage.SaveConfig(path, cfg); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	err := Run(context.Background(), []string{"set-goal", "New", "Goal"}, Env{
		ConfigPath: path,
		Out:        &out,
		Err:        &errOut,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	updated, err := storage.LoadConfig(path)
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}
	if updated.Goal != "New Goal" {
		t.Fatalf("expected goal to update, got %q", updated.Goal)
	}
}

func TestRunAddTask(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cfg := storage.Config{Goal: "Ship"}
	path := t.TempDir() + "/goalterm.json"
	if err := storage.SaveConfig(path, cfg); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	err := Run(context.Background(), []string{"add-task", "Write", "tests"}, Env{
		ConfigPath: path,
		Out:        &out,
		Err:        &errOut,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	updated, err := storage.LoadConfig(path)
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}
	if len(updated.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(updated.Tasks))
	}
	if updated.Tasks[0].Title != "Write tests" {
		t.Fatalf("expected task title to update, got %q", updated.Tasks[0].Title)
	}
}

func TestRunRemoveTask(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cfg := storage.Config{
		Goal: "Ship",
		Tasks: []goal.Task{
			{Title: "First", Status: goal.StatusTodo},
			{Title: "Second", Status: goal.StatusTodo},
		},
	}
	path := t.TempDir() + "/goalterm.json"
	if err := storage.SaveConfig(path, cfg); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	err := Run(context.Background(), []string{"remove", "1"}, Env{
		ConfigPath: path,
		Out:        &out,
		Err:        &errOut,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	updated, err := storage.LoadConfig(path)
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}
	if len(updated.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(updated.Tasks))
	}
	if updated.Tasks[0].Title != "Second" {
		t.Fatalf("expected remaining task to be second, got %q", updated.Tasks[0].Title)
	}
}

func TestRunCompleteTask(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cfg := storage.Config{
		Goal: "Ship",
		Tasks: []goal.Task{
			{Title: "First", Status: goal.StatusTodo},
		},
	}
	path := t.TempDir() + "/goalterm.json"
	if err := storage.SaveConfig(path, cfg); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	err := Run(context.Background(), []string{"complete", "1"}, Env{
		ConfigPath: path,
		Out:        &out,
		Err:        &errOut,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	updated, err := storage.LoadConfig(path)
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}
	if updated.Tasks[0].Status != goal.StatusDone {
		t.Fatalf("expected task to be done, got %q", updated.Tasks[0].Status)
	}
}

func TestRunHelpOutputsUsage(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cfg := storage.Config{Goal: "Ship"}
	path := t.TempDir() + "/goalterm.json"
	if err := storage.SaveConfig(path, cfg); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	err := Run(context.Background(), []string{"help"}, Env{
		ConfigPath: path,
		Out:        &out,
		Err:        &errOut,
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !bytes.Contains(out.Bytes(), []byte("Usage:")) {
		t.Fatalf("expected usage in output")
	}
}
