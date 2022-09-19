package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTailscaleDevice(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_device",
		Description: "A Tailscale Device represents the devices under a tailnet.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleDevices,
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
				Name:        "name",
				Description: "The name of the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of the device.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listTailscaleDevices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_device.listTailscaleDevices", "connection_error", err)
		return nil, err
	}

	// req := tailscale.GetDevicesArgs{}

	// devices, err := conn.AlertContact.GetAlertContacts(params)
	devices, err := client.Devices(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range devices {
		d.StreamListItem(ctx, item)
		// return nil
	}

	return nil, nil
}
