package suggest

import (
	"context"

	"github.com/murraycode/goal-term/internal/storage"
)

type Suggester interface {
	Suggest(ctx context.Context, cfg storage.Config) (string, error)
}
