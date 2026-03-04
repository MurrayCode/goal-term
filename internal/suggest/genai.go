package suggest

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/murraycode/goal-term/internal/storage"
	"google.golang.org/genai"
)

type GenAISuggester struct {
	Client *genai.Client
}

func NewGenAISuggester(apiKey string) (*GenAISuggester, error) {
	timeout := 5 * time.Minute
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:      apiKey,
		Backend:     genai.BackendGeminiAPI,
		HTTPOptions: genai.HTTPOptions{Timeout: &timeout},
	})
	if err != nil {
		return nil, err
	}

	return &GenAISuggester{Client: client}, nil
}

func (s *GenAISuggester) Suggest(ctx context.Context, cfg storage.Config) (string, error) {
	if s == nil || s.Client == nil {
		return "", nil
	}

	prompt := buildPrompt(cfg)
	resp, err := s.Client.Models.GenerateContent(ctx, "gemini-2.5-flash", genai.Text(prompt), nil)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(resp.Text()), nil
}

func buildPrompt(cfg storage.Config) string {
	var b strings.Builder
	b.WriteString("Analyze the goal and tasks. Provide 3-5 concise suggestions on what to look at next.\n\n")
	b.WriteString("Goal: ")
	b.WriteString(cfg.Goal)
	b.WriteString("\n")

	if len(cfg.Tasks) == 0 {
		b.WriteString("Tasks: none\n")
		return b.String()
	}

	b.WriteString("Tasks:\n")
	for i, task := range cfg.Tasks {
		b.WriteString(strconv.Itoa(i + 1))
		b.WriteString(". ")
		b.WriteString(task.Title)
		b.WriteString(" (")
		b.WriteString(string(task.Status))
		b.WriteString(")\n")
	}

	return b.String()
}
