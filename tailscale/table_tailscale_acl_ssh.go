package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tablelistTailscaleAclSsh(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_acl_ssh",
		Description: "Tailscale Acl Ssh.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleAclSsh,
		},
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
		},
	}
}

func listTailscaleAclSsh(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// retrieves the ACL that is currently set for the given tailnet.
	acl, err := getTailscaleAcl(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleAclSsh", "connection_error", err)
		return nil, err
	}

	for _, element := range acl.SSH {
		d.StreamListItem(ctx, element)
	}
	return nil, nil
}
