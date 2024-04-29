package tailscale

import (
	"context"

	"github.com/tailscale/tailscale-client-go/tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
				Description: "If default regions should be omitted.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "regions",
				Description: "List of regions.",
				Type:        proto.ColumnType_JSON,
			},
		}),
	}
}

func listTailscaleACLDERPMap(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// retrieves the ACL that is currently set for the given tailnet.
	data, err := getTailscaleAcl(ctx, d, h)
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
