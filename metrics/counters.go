package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	FeatureToggleErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "feature_toggle_errors",
		Help: "counter for total number of errors from feature toggle server",
	})

	FeatureToggleQueries = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "feature_toggle_queries",
		Help: "counter for query of feature toggles",
	}, []string{"name"})
)
