package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableTailscaleAclAutoApprovers(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_acl_auto_approvers",
		Description: "Tailscale ACL Auto Approvers.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleAclAutoApprovers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "routes",
				Description: ".",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "exit_nodes",
				Description: ".",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

func listTailscaleAclAutoApprovers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// retrieves the ACL that is currently set for the given tailnet.
	acl, err := getTailscaleAcl(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleAclSsh", "connection_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, acl.AutoApprovers)
	return nil, nil
}
