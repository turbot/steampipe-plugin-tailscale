package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func getTailscaleAcl(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.getTailscaleAcl", "connection_error", err)
		return nil, err
	}

	acl, err := client.ACL(ctx)
	if err != nil {
		return nil, err
	}
	return acl, nil
}
