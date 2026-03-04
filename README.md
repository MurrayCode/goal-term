# goal-term

Tiny Go starter for storing goals and tasks.

## Quick start

```bash
go run ./cmd/goal-term
```

## Commands

```bash
goal-term list
goal-term set-goal "Ship v1"
goal-term add-task "Write CLI commands"
goal-term complete 2
goal-term remove 1
```

## Gemini suggestions

Set a Gemini API key to enable suggestions on startup:

```bash
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

Add this to your `~/.bashrc`:

```bash
# goal-term
if [[ $- == *i* ]]; then
  GOALTERM_CONFIG="$HOME/Development/goal-term/goalterm.json" \
    goal-term list
fi
```

Add this to your `~/.zshrc`:

```zsh
# goal-term
if [[ $- == *i* ]]; then
  GOALTERM_CONFIG="$HOME/Development/goal-term/goalterm.json" \
    goal-term list
fi
```

## Config file

The goal and tasks are stored in `goalterm.json` at the project root.

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
- `internal/goal`: goal and task domain types
- `internal/storage`: config file loader/saver
```
