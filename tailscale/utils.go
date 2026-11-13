package tailscale

import (
	"context"
	"strings"

	"github.com/tailscale/tailscale-client-go/tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v6/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

var getTailscaleAclMemoized = plugin.HydrateFunc(getTailscaleAclUncached).Memoize(memoize.WithCacheKeyFunction(getTailscaleAclCacheKey))

func getTailscaleAclCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getTailscaleAcl"
	return key, nil
}

func getTailscaleAcl(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	tailscaleInfo, err := getTailscaleAclMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}

	return tailscaleInfo, nil
}

func getTailscaleAclUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheKey := "getTailscaleAcl"
	var err error
	var tailscaleInfo *tailscale.ACL

	if cachedData, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		tailscaleInfo = cachedData.(*tailscale.ACL)
	} else {
		client, err := connect(ctx, d)
		if err != nil {
			plugin.Logger(ctx).Error("tailscale_acl.getTailscaleAcl", "connection_error", err)
			return nil, err
		}

		tailscaleInfo, err = client.ACL(ctx)
		if err != nil {
			return nil, err
		}
	}

	return tailscaleInfo, err
}

func isNotFoundError(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
	return strings.Contains(err.Error(), "404")
}
