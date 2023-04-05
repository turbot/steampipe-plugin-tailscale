package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// column definitions for the common columns
func commonColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "tailnet_name",
			Description: "The name of your tailnet.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getTailscaleTailnet,
			Transform:   transform.FromValue(),
		},
	}
}

func defaultColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumns()...)
}

func getTailscaleTailnet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	if d.EqualsQualString("tailnet_name") != "" && d.EqualsQualString("tailnet_name") != *config.TailnetName {
		return nil, nil
	}

	if config.TailnetName != nil {
		return config.TailnetName, nil
	}
	return nil, nil
}
