# goal-term

Tiny Go starter for storing goals and tasks. Provides suggestions on progressing goals & tasks using Gemini.

![Example](images/example)

## Quick start

```bash
export GOALTERM_CONFIG="$HOME/.config/goal-term/goalterm.json"
go run ./cmd/goal-term
```

## Commands

```bash
export GOALTERM_CONFIG="$HOME/.config/goal-term/goalterm.json"
goal-term list
goal-term set-goal "Ship v1"
goal-term add-task "Write CLI commands"
goal-term complete 2
goal-term remove 1
```

## Gemini suggestions

Set a Gemini API key to enable suggestions on startup:

```bash
export GOALTERM_CONFIG="$HOME/.config/goal-term/goalterm.json"
export GOOGLE_API_KEY="your-api-key"
```

## Build and install

Build a local binary:

```bash
go build -o bin/goal-term ./cmd/goal-term
```

Install using the module path:

```bash
go install github.com/murraycode/goal-term/cmd/goal-term@latest
```

Make sure your Go bin directory is on PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Run on new terminals

You are responsible for choosing where the config lives and exporting it in your shell startup.

Add this to your `~/.bashrc`:

```bash
# goal-term
export GOALTERM_CONFIG="$HOME/.config/goal-term/goalterm.json"
if [[ $- == *i* ]]; then
  goal-term list
fi
```

Add this to your `~/.zshrc`:

```zsh
# goal-term
export GOALTERM_CONFIG="$HOME/.config/goal-term/goalterm.json"
if [[ $- == *i* ]]; then
  goal-term list
fi
```

## Config file

The goal and tasks are stored in the file pointed to by `GOALTERM_CONFIG`.

Example:

```json
{
  "goal": "Learn Go",
  "tasks": [
    {"title": "Read the Go Tour", "status": "todo"},
    {"title": "Build a small CLI", "status": "todo"}
  ]
}
```

## Project layout

- `cmd/goal-term`: entrypoint
- `internal/app`: orchestration and command handling
- `internal/cli`: typed command parsing
- `internal/goal`: goal and task domain types
- `internal/storage`: config file loader/saver
- `internal/suggest`: Gemini suggestions interface + implementation
- `internal/ui`: output formatting
```
