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
		Description: "Tailscale ACL Entry.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleAclEntry,
		},
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
		},
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
