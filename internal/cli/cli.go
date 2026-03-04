package cli

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CommandType string

const (
	CommandList     CommandType = "list"
	CommandSetGoal  CommandType = "set-goal"
	CommandAddTask  CommandType = "add-task"
	CommandComplete CommandType = "complete"
	CommandRemove   CommandType = "remove"
	CommandHelp     CommandType = "help"
)

type Command struct {
	Type  CommandType
	Title string
	Index int
}

var ErrUsage = errors.New("usage")

func Parse(args []string) (Command, error) {
	if len(args) == 0 {
		return Command{Type: CommandList}, nil
	}

	switch args[0] {
	case "list":
		return Command{Type: CommandList}, nil
	case "set-goal":
		title := strings.TrimSpace(strings.Join(args[1:], " "))
		if title == "" {
			return Command{}, usageError("set-goal requires a goal title")
		}
		return Command{Type: CommandSetGoal, Title: title}, nil
	case "add-task":
		title := strings.TrimSpace(strings.Join(args[1:], " "))
		if title == "" {
			return Command{}, usageError("add-task requires a task title")
		}
		return Command{Type: CommandAddTask, Title: title}, nil
	case "complete":
		index, err := parseIndex(args)
		if err != nil {
			return Command{}, err
		}
		return Command{Type: CommandComplete, Index: index}, nil
	case "remove":
		index, err := parseIndex(args)
		if err != nil {
			return Command{}, err
		}
		return Command{Type: CommandRemove, Index: index}, nil
	case "help", "-h", "--help":
		return Command{Type: CommandHelp}, nil
	default:
		return Command{}, usageError("unknown command: " + args[0])
	}
}

func Usage(w io.Writer, errMsg string) {
	if errMsg != "" {
		fmt.Fprintln(w, errMsg)
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  goal-term list")
	fmt.Fprintln(w, "  goal-term set-goal <goal title>")
	fmt.Fprintln(w, "  goal-term add-task <task title>")
	fmt.Fprintln(w, "  goal-term complete <task number>")
	fmt.Fprintln(w, "  goal-term remove <task number>")
}

func parseIndex(args []string) (int, error) {
	if len(args) < 2 {
		return 0, usageError("command requires a task number")
	}

	value, err := strconv.Atoi(args[1])
	if err != nil || value < 1 {
		return 0, usageError("task number must be a positive integer")
	}

	return value - 1, nil
}

func usageError(msg string) error {
	return fmt.Errorf("%w: %s", ErrUsage, msg)
}
