package feature_flags

import (
	"fmt"

	"github.com/go-coldbrew/feature-flags/config"
	"github.com/go-coldbrew/feature-flags/engines"
	"github.com/go-coldbrew/feature-flags/engines/unleash"
)

var client engines.FeatureFlag

const (
	EngineUnleash = "unleash"
)

// Initialize initializes an engines.FeatureFlag instance, return error if failed
func Initialize(appName string, cfg config.Config) error {
	var err error
	switch cfg.FeatureFlagEngine {
	case EngineUnleash:
		client, err = unleash.Initialize(appName, cfg.UnleashConfig)
		return err
	default:
		return fmt.Errorf("unsupported feature flag engine: %s", cfg.FeatureFlagEngine)
	}
}

// IsEnabled check if a feature flag is enabled, returns false if client is not initialized
func IsEnabled(name string, ctx engines.Context) bool {
	if client == nil {
		return false
	}

	return client.IsEnabled(name, ctx)
}

// GetVariant get variant for a feature flag, returns disabled variant if client is not initialised
func GetVariant(name string, ctx engines.Context) *engines.Variant {
	if client == nil {
		return engines.DisabledVariant
	}

	return client.GetVariant(name, ctx)
}
