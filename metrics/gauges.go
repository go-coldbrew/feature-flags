package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	FeatureToggleRepositoryReady = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "feature_toggle_repository_ready",
		Help: "gauge for feature toggle repository ready event",
	})
)
