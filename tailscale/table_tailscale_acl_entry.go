package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	// "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTailscaleAclEntry(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_acl_entry",
		Description: "Tailscale Acl Entry.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleAclEntry,
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
				Name:        "ports",
				Description: ".",
				Type:        proto.ColumnType_JSON,
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
				Name:        "protocol",
				Description: ".",
				Type:        proto.ColumnType_STRING,
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

func listTailscaleAclEntry(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleAclEntry", "connection_error", err)
		return nil, err
	}
	acl, err := client.ACL(ctx)
	if err != nil {
		return nil, err
	}
	for _, element := range acl.ACLs {
		d.StreamListItem(ctx, element)
	}
	return nil, nil

}
