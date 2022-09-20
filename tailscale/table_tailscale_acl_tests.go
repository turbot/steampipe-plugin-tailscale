package tailscale

import (
	"context"

	"github.com/tailscale/tailscale-client-go/tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableTailscaleAclTest(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_acl_test",
		Description: "Tailscale ACL Test.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleACLTests,
		},
		Columns: defaultColumns([]*plugin.Column{
			{
				Name:        "source",
				Description: ".",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user",
				Description: ".",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "accept",
				Description: ".",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "allow",
				Description: ".",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "deny",
				Description: ".",
				Type:        proto.ColumnType_JSON,
			},
		}),
	}
}

func listTailscaleACLTests(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// retrieves the ACL that is currently set for the given tailnet.
	getTailscaleAclCached := plugin.HydrateFunc(getTailscaleAcl).WithCache()
	data, err := getTailscaleAclCached(ctx, d, h)
	if data == nil {
		return nil, nil
	}
	acl := data.(*tailscale.ACL)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleACLTests", "connection_error", err)
		return nil, err
	}

	for _, element := range acl.Tests {
		d.StreamListItem(ctx, element)
	}
	return nil, nil
}
