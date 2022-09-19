package tailscale

import (
	"context"

	"github.com/tailscale/tailscale-client-go/tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func listTailscaleAcl(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (*tailscale.ACL, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleAclSsh", "connection_error", err)
		return nil, err
	}
	acl, err := client.ACL(ctx)
	if err != nil {
		return nil, err
	}
	return acl, nil
}
