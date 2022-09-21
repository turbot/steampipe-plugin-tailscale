package tailscale

import (
	"context"

	"github.com/tailscale/tailscale-client-go/tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableTailscaleAclEntry(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_acl_entry",
		Description: "Tailscale ACL entry.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleAclEntry,
		},
		Columns: defaultColumns([]*plugin.Column{
			{
				Name:        "action",
				Description: "Action defined for a device or a network.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ports",
				Description: "The list of ports to apply the action on.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "users",
				Description: "The list of users to apply an action on.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "source",
				Description: "The list of source IP addresses for the connection.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "destination",
				Description: "The list of destination IP addresses for the connection.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "protocol",
				Description: "The protocol of the connection.",
				Type:        proto.ColumnType_STRING,
			},
		}),
	}
}

func listTailscaleAclEntry(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// retrieves the ACL that is currently set for the given tailnet.
	getTailscaleAclCached := plugin.HydrateFunc(getTailscaleAcl).WithCache()
	data, err := getTailscaleAclCached(ctx, d, h)
	if data == nil {
		return nil, nil
	}
	acl := data.(*tailscale.ACL)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleAclSsh", "connection_error", err)
		return nil, err
	}

	for _, element := range acl.ACLs {
		d.StreamListItem(ctx, element)
	}
	return nil, nil
}
