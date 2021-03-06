package unleash

import (
	"context"
	"net/http"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/Unleash/unleash-client-go/v3/api"
	unleashCtx "github.com/Unleash/unleash-client-go/v3/context"
	"github.com/go-coldbrew/feature-flags/config"
	"github.com/go-coldbrew/feature-flags/engines"
	"github.com/go-coldbrew/log"
)

type unleashClient struct {
}

func Initialize(appName string, cfg config.UnleashConfig) (engines.FeatureFlag, error) {
	if cfg.UnleashUrl == "" {
		log.Info(context.Background(), "UNLEASH_URL is not configured, no feature flags in action")
		return nil, nil
	}

	err := unleash.Initialize(
		unleash.WithListener(&LogListener{}),
		unleash.WithAppName(appName),
		unleash.WithUrl(cfg.UnleashUrl),
		unleash.WithCustomHeaders(http.Header{"Authorization": {cfg.UnleashToken}}))
	if err != nil {
		log.Error(context.Background(), "Failed initializing Unleash client", err)
		return nil, err
	}
	return &unleashClient{}, nil
}

func mapToUnleashContext(ctx engines.Context) unleashCtx.Context {
	return unleashCtx.Context{
		UserId:        ctx.UserId,
		SessionId:     ctx.SessionId,
		RemoteAddress: ctx.RemoteAddress,
		Environment:   ctx.Environment,
		AppName:       ctx.AppName,
		Properties:    ctx.Properties,
	}
}

func mapFromUnleashVariant(variant *api.Variant) *engines.Variant {
	payload := engines.Payload{
		Type:  variant.Payload.Type,
		Value: variant.Payload.Value,
	}
	return &engines.Variant{
		Name:    variant.Name,
		Payload: payload,
		Enabled: variant.Enabled,
	}
}

func (c *unleashClient) IsEnabled(name string, ctx engines.Context) bool {
	return unleash.IsEnabled(name, unleash.WithContext(mapToUnleashContext(ctx)))
}

func (c *unleashClient) GetVariant(name string, ctx engines.Context) *engines.Variant {
	variant := unleash.GetVariant(name, unleash.WithVariantContext(mapToUnleashContext(ctx)))
	return mapFromUnleashVariant(variant)
}
