package tailscale

import (
	"context"

	"github.com/tailscale/tailscale-client-go/tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableTailscaleAclDERPMap(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_acl_derp_map",
		Description: "Tailscale ACL DERP Map.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleACLDERPMap,
		},
		Columns: defaultColumns([]*plugin.Column{
			{
				Name:        "omit_default_regions",
				Description: ".",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "regions",
				Description: ".",
				Type:        proto.ColumnType_JSON,
			},
		}),
	}
}

func listTailscaleACLDERPMap(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// retrieves the ACL that is currently set for the given tailnet.
	getTailscaleAclCached := plugin.HydrateFunc(getTailscaleAcl).WithCache()
	data, err := getTailscaleAclCached(ctx, d, h)
	if data == nil {
		return nil, nil
	}
	acl := data.(*tailscale.ACL)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_acl.listTailscaleACLDERPMap", "connection_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, acl.DERPMap)
	return nil, nil
}
