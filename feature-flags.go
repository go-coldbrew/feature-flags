package feature_flags

import (
	"fmt"

	"github.com/go-coldbrew/feature-flags/config"
	"github.com/go-coldbrew/feature-flags/engines"
	"github.com/go-coldbrew/feature-flags/engines/unleash"
)

const (
	EngineUnleash = "unleash"
)

// New creates a new engines.FeatureFlag instance
func New(appName string, cfg config.Config) (engines.FeatureFlag, error) {
	switch cfg.FeatureFlagEngine {
	case EngineUnleash:
		return unleash.Initialize(appName, cfg.UnleashConfig)
	default:
		return nil, fmt.Errorf("unsupported feature flag engine: %s", cfg.FeatureFlagEngine)
	}
}
