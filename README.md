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

## Show on new terminal (zsh)

Add this to your `~/.zshrc`:

```zsh
# goal-term
(cd /path/to/goal-term && go run ./cmd/goal-term)
```

## Show on new terminal (bash)

Add this to your `~/.bashrc`:

```bash
# goal-term
if [[ $- == *i* ]]; then
  GOALTERM_CONFIG="$HOME/Development/goal-term/goalterm.json" \
    go run "$HOME/Development/goal-term/cmd/goal-term"
fi
```

## Project layout

- `cmd/goal-term`: entrypoint
- `internal/goal`: goal and task domain types
- `internal/storage`: config file loader/saver
```
