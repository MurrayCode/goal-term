package cli

import (
	"errors"
	"testing"
)

func TestParseListDefault(t *testing.T) {
	cmd, err := Parse(nil)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if cmd.Type != CommandList {
		t.Fatalf("expected list command, got %q", cmd.Type)
	}
}

func TestParseSetGoal(t *testing.T) {
	cmd, err := Parse([]string{"set-goal", "Ship", "v1"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if cmd.Type != CommandSetGoal {
		t.Fatalf("expected set-goal command, got %q", cmd.Type)
	}
	if cmd.Title != "Ship v1" {
		t.Fatalf("expected title %q, got %q", "Ship v1", cmd.Title)
	}
}

func TestParseCompleteInvalid(t *testing.T) {
	_, err := Parse([]string{"complete", "0"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, ErrUsage) {
		t.Fatalf("expected usage error, got %v", err)
	}
}

func TestParseUnknownCommand(t *testing.T) {
	_, err := Parse([]string{"whoami"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, ErrUsage) {
		t.Fatalf("expected usage error, got %v", err)
	}
}
