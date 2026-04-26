package app

import (
	"github.com/TehranTime/tehtime-core/internal/config"
	"go.uber.org/zap"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	// For now, use zap's production config; you can switch to development based on cfg.Debug
	if cfg != nil && cfg.Debug {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}
