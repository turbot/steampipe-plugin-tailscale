package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	// "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tablelistTailscaleAclSsh(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_acl_ssh",
		Description: "Tailscale Acl Ssh.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleAclSsh,
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
				Name:        "action",
				Description: ".",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "users",
				Description: ".",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "source",
				Description: ".",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "destination",
				Description: ".",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "check_period",
				Description: ".",
				Type:        proto.ColumnType_TIMESTAMP,
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

func listTailscaleAclSsh(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleAclSsh", "connection_error", err)
		return nil, err
	}
	acl, err := client.ACL(ctx)
	if err != nil {
		return nil, err
	}
	for _, element := range acl.SSH {
		d.StreamListItem(ctx, element)
	}
	return nil, nil
}
