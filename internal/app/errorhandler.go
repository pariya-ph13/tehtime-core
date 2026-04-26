package app

import (
	"github.com/TehranTime/tehtime-core/internal/config"
)

// newSentry is currently a no-op placeholder. Wire a real error handler here later.
func newSentry(cfg *config.Config) error {
	_ = cfg
	return nil
}
