package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/murraycode/goal-term/internal/goal"
)

type Config struct {
	Goal  string      `json:"goal"`
	Tasks []goal.Task `json:"tasks"`
}

var (
	ErrConfigNotFound   = errors.New("config not found")
	ErrConfigUnreadable = errors.New("config unreadable")
)

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Config{}, ErrConfigNotFound
		}
		return Config{}, ErrConfigUnreadable
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func SaveConfig(path string, cfg Config) error {
	if strings.TrimSpace(cfg.Goal) == "" {
		return errors.New("goal must not be empty")
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}
