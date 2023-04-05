package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTailscaleTailnetKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_tailnet_key",
		Description: "A Tailscale tailnet key represents the keys for a tailnet.",
		Get: &plugin.GetConfig{
			Hydrate: getTailscaleTailnetKey,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "id",
					Require: plugin.Required,
				},
			},
		},
		Columns: defaultColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "An unique identifier of the tailnet key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "key",
				Description: "Key information.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created",
				Description: "Device creation time.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "expires",
				Description: "Device expiry time.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "capabilities",
				Description: "The list of device capabilities.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Capabilities"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Key"),
			},
		}),
	}
}

//// LIST FUNCTION

func getTailscaleTailnetKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleTailnetKey", "connection_error", err)
		return nil, err
	}

	id := d.EqualsQuals["id"].GetStringValue()
	if id == "" {
		return nil, nil
	}

	key, err := client.GetKey(ctx, id)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleTailnetKey", "api_error", err)
		return nil, err
	}

	return key, nil
}
