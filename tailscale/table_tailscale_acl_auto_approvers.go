package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	// "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTailscaleAclAutoApprovers(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_acl_auto_approvers",
		Description: "Tailscale Acl Auto Approvers.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleAclAutoApprovers,
			// KeyColumns: []*plugin.KeyColumn{
			// 	{
			// 		Name:    "name",
			// 		Require: plugin.Optional,
			// 	},
			// },
		},
		// Get: &plugin.GetConfig{
		// 	Hydrate:    gettailscaleVendor,
		// 	KeyColumns: plugin.SingleColumn("id"),
		// },
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
			// Steampipe standard columns
			// {
			// 	Name:        "title",
			// 	Description: "Title of the resource.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Transform:   transform.FromField("Name"),
			// },
		},
	}
}

func listTailscaleAclAutoApprovers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	acl, err := listTailscaleAcl(ctx, d, h)

	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleAclSsh", "connection_error", err)
		return nil, err
	} else {
		d.StreamListItem(ctx, acl.AutoApprovers)
		return nil, nil
	}
}
