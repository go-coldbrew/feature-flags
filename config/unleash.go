package config

type UnleashConfig struct {
	UnleashUrl   string `envconfig:"UNLEASH_URL"`
	UnleashToken string `envconfig:"UNLEASH_TOKEN"`
	UnleashRetryIntervalSeconds int `envconfig:"UNLEASH_RETRY_INTERVAL_SECONDS" default:"60"`
}
