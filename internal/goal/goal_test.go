package goal

import "testing"

func TestNewSeedsTodoTasks(t *testing.T) {
	goal := New("Ship v1", []string{"Plan", "Execute"})
	if goal.Title != "Ship v1" {
		t.Fatalf("expected goal title, got %q", goal.Title)
	}
	if len(goal.Tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(goal.Tasks))
	}
	if goal.Tasks[0].Status != StatusTodo {
		t.Fatalf("expected first task todo, got %q", goal.Tasks[0].Status)
	}
	if goal.Tasks[1].Status != StatusTodo {
		t.Fatalf("expected second task todo, got %q", goal.Tasks[1].Status)
	}
}
