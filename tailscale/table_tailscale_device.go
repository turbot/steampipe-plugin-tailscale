package tailscale

import (
	"context"

	"github.com/tailscale/tailscale-client-go/tailscale"
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
		},
		Columns: defaultColumns([]*plugin.Column{
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
			{
				Name:        "authorized",
				Description: "Whether the device is authorized.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "blocks_incoming_connections",
				Description: "Whether the device blocks incoming connections.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "client_version",
				Description: "Version of the client.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created",
				Description: "Device creation time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Created.Time"),
			},
			{
				Name:        "expires",
				Description: "Device expiry time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Expires.Time"),
			},
			{
				Name:        "hostname",
				Description: "Name of the host.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_external",
				Description: "Whether the device is external.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "key_expiry_disabled",
				Description: "Whether the key expiration is disabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "last_seen",
				Description: "Device last active time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("LastSeen.Time"),
			},
			{
				Name:        "machine_key",
				Description: "Machine key of the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "node_key",
				Description: "Node key of the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "os",
				Description: "OS information of the device.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OS"),
			},
			{
				Name:        "update_available",
				Description: "Whether an update is available.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "user",
				Description: "Name of the owner of the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "addresses",
				Description: "The list of device's IPs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "device_subnet_routes",
				Description: "A list of the device subnet routes.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getTailscaleDeviceSubnetRoutes,
			},
			{
				Name:        "tags",
				Description: "The tags applied to the device.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listTailscaleDevices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("tailscale_device.listTailscaleDevices", "connection_error", err)
		return nil, err
	}

	devices, err := client.Devices(ctx)
	if err != nil {
		logger.Error("tailscale_device.listTailscaleDevices", "api_error", err)
		return nil, err
	}
	for _, item := range devices {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

func getTailscaleDeviceSubnetRoutes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	device := h.Item.(tailscale.Device)
	id := device.ID

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleDeviceRoutes", "connection_error", err)
		return nil, err
	}

	deviceRoutes, err := client.DeviceSubnetRoutes(ctx, id)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleDeviceRoutes", "api_error", err)
		return nil, err
	}

	return deviceRoutes, nil
}
