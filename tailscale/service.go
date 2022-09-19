package tailscale

import (
	"context"
	"errors"
	"os"

	"github.com/tailscale/tailscale-client-go/tailscale"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*tailscale.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "tailscale"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*tailscale.Client), nil
	}

	// Default to using env vars (#2)
	apiKey := os.Getenv("TAILSCALE_API_KEY")
	tailnetName := os.Getenv("TAILSCALE_TAILNET_NAME")

	// But prefer the config (#1)
	tailscaleConfig := GetConfig(d.Connection)
	if tailscaleConfig.APIKey != nil {
		apiKey = *tailscaleConfig.APIKey
	}
	if tailscaleConfig.TailnetName != nil {
		tailnetName = *tailscaleConfig.TailnetName
	}

	if apiKey == "" || tailnetName == "" {
		// Credentials not set
		return nil, errors.New("api_key and tailnet must be configured")
	}

	// Configure to automatically wait 1 sec between requests, per Zoom API requirements
	conn, err := tailscale.NewClient(apiKey, tailnetName)

	if err != nil {
		d.ConnectionManager.Cache.Set(cacheKey, conn)
	}

	return conn, nil
}
