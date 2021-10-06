package config

type Config struct {
	FeatureFlagEngine string `envconfig:"FEATURE_FLAG_ENGINE" default:"unleash"`
	UnleashConfig
}
