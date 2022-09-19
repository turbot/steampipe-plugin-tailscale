package tailscale

import (
	"context"
	"fmt"
	"os"

	"https://github.com/pulumi/pulumi-tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

// getSessionConfig :: returns tailscale client to perform API requests
func getSessionConfig(ctx context.Context, d *plugin.QueryData) (*tailscale.Client, error) {
	// Load clientOptions from cache
	sessionCacheKey := "tailscale.clientoption"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*tailscale.Client), nil
	}

	// Get tailscale config
	tailscaleConfig := GetConfig(d.Connection)

	// Get the authorization token
	token := os.Getenv("TAILSCALE_TOKEN")
	if tailscaleConfig.Token != nil {
		token = *tailscaleConfig.Token
	}

	// No creds
	if token == "" {
		return nil, fmt.Errorf("API KEY must be configured")
	}

	// Create client
	client := tailscale.NewClient(token)

	// save clientOptions in cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}
