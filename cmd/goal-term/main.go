package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/murraycode/goal-term/internal/app"
	"github.com/murraycode/goal-term/internal/suggest"
)

func main() {
	configPath := strings.TrimSpace(os.Getenv("GOALTERM_CONFIG"))
	if configPath == "" {
		fmt.Fprintln(os.Stderr, "GOALTERM_CONFIG must be set to a config file path")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var suggester suggest.Suggester
	if apiKey := strings.TrimSpace(os.Getenv("GOOGLE_API_KEY")); apiKey != "" {
		genaiSuggester, err := suggest.NewGenAISuggester(apiKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to initialize genai client: %v\n", err)
		} else {
			suggester = genaiSuggester
		}
	}

	err := app.Run(ctx, os.Args[1:], app.Env{
		ConfigPath: configPath,
		Out:        os.Stdout,
		Err:        os.Stderr,
		Suggester:  suggester,
	})
	if err != nil {
		os.Exit(1)
	}
}
