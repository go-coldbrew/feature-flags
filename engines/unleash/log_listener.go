package unleash

import (
	"context"

	"github.com/go-coldbrew/feature-flags/metrics"
	"github.com/go-coldbrew/log"
)

// LogListener is an implementation of the listener interfaces that simply logs everything
type LogListener struct{}

// OnError logs errors.
func (l LogListener) OnError(err error) {
	metrics.FeatureToggleErrors.Inc()
	log.Error(context.Background(), err)
}

// OnWarning logs warning.
func (l LogListener) OnWarning(warning error) {
	log.Warn(context.Background(), warning)
}

// OnReady logs when the repository is ready.
func (l LogListener) OnReady() {
	metrics.FeatureToggleRepositoryReady.Set(1)
	log.Info(context.Background(), "Unleash repository READY")
}

// OnCount emits metrics when the feature is queried.
func (l LogListener) OnCount(name string, enabled bool) {
	metrics.FeatureToggleQueries.WithLabelValues(name).Inc()
	log.Debug(context.Background(), "msg", "Queried feature toggle", "toggle", name, "enabled", enabled)
}
