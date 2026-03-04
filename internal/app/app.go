package app

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/murraycode/goal-term/internal/cli"
	"github.com/murraycode/goal-term/internal/goal"
	"github.com/murraycode/goal-term/internal/storage"
	"github.com/murraycode/goal-term/internal/suggest"
	"github.com/murraycode/goal-term/internal/ui"
)

type Env struct {
	ConfigPath string
	Out        io.Writer
	Err        io.Writer
	Suggester  suggest.Suggester
}

func Run(ctx context.Context, args []string, env Env) error {
	if env.Out == nil {
		env.Out = io.Discard
	}
	if env.Err == nil {
		env.Err = io.Discard
	}

	cfg, err := loadOrSeedConfig(env.ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	cmd, err := cli.Parse(args)
	if err != nil {
		if errors.Is(err, cli.ErrUsage) {
			cli.Usage(env.Err, err.Error())
			return err
		}
		return err
	}

	switch cmd.Type {
	case cli.CommandList:
		ui.PrintGoal(env.Out, cfg)
		return maybeSuggest(ctx, env, cfg)
	case cli.CommandSetGoal:
		cfg.Goal = cmd.Title
		if err := storage.SaveConfig(env.ConfigPath, cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		ui.PrintGoal(env.Out, cfg)
	case cli.CommandAddTask:
		cfg.Tasks = append(cfg.Tasks, goal.Task{Title: cmd.Title, Status: goal.StatusTodo})
		if err := storage.SaveConfig(env.ConfigPath, cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		ui.PrintGoal(env.Out, cfg)
	case cli.CommandComplete:
		if cmd.Index >= len(cfg.Tasks) {
			cli.Usage(env.Err, "task number out of range")
			return cli.ErrUsage
		}
		cfg.Tasks[cmd.Index].Status = goal.StatusDone
		if err := storage.SaveConfig(env.ConfigPath, cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		ui.PrintGoal(env.Out, cfg)
	case cli.CommandRemove:
		if cmd.Index >= len(cfg.Tasks) {
			cli.Usage(env.Err, "task number out of range")
			return cli.ErrUsage
		}
		cfg.Tasks = append(cfg.Tasks[:cmd.Index], cfg.Tasks[cmd.Index+1:]...)
		if err := storage.SaveConfig(env.ConfigPath, cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
		ui.PrintGoal(env.Out, cfg)
	case cli.CommandHelp:
		cli.Usage(env.Out, "")
		return nil
	default:
		return fmt.Errorf("unknown command: %s", cmd.Type)
	}

	return nil
}

func loadOrSeedConfig(path string) (storage.Config, error) {
	cfg, err := storage.LoadConfig(path)
	if err == nil {
		return cfg, nil
	}
	if !errors.Is(err, storage.ErrConfigNotFound) && !errors.Is(err, storage.ErrConfigUnreadable) {
		return storage.Config{}, err
	}
	if errors.Is(err, storage.ErrConfigUnreadable) {
		return storage.Config{}, err
	}

	seed := storage.Config{
		Goal: "Learn Go",
		Tasks: []goal.Task{
			{Title: "Read the Go Tour", Status: goal.StatusTodo},
			{Title: "Build a small CLI", Status: goal.StatusTodo},
		},
	}
	if err := storage.SaveConfig(path, seed); err != nil {
		return storage.Config{}, err
	}

	return seed, nil
}

func maybeSuggest(ctx context.Context, env Env, cfg storage.Config) error {
	if env.Suggester == nil {
		return nil
	}

	suggestions, err := env.Suggester.Suggest(ctx, cfg)
	if err != nil {
		fmt.Fprintf(env.Err, "failed to get suggestions: %v\n", err)
		return nil
	}

	ui.PrintSuggestions(env.Out, suggestions)
	return nil
}
