package tailscale

import (
	"context"

	"github.com/tailscale/tailscale-client-go/tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableTailscaleAclAutoApprover(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_acl_auto_approver",
		Description: "Tailscale ACL Auto Approvers.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleAclAutoApprovers,
		},
		Columns: defaultColumns([]*plugin.Column{
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
		}),
	}
}

func listTailscaleAclAutoApprovers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// retrieves the ACL that is currently set for the given tailnet.
	getTailscaleAclCached := plugin.HydrateFunc(getTailscaleAcl).WithCache()
	data, err := getTailscaleAclCached(ctx, d, h)
	if data == nil {
		return nil, nil
	}
	acl := data.(*tailscale.ACL)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleAclAutoApprovers", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, acl.AutoApprovers)
	return nil, nil
}
