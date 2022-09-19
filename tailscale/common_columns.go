package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

// column definitions for the common columns
func commonColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "tailnet_name",
			Description: "The name of your private network.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromValue(),
		},
	}
}

func defaultColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumns()...)
}

// get columns which are returned with all tables: region, partition and account
func getTailscaleTailnet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	if config.Tailnet != nil {
		return config.Tailnet, nil
	}
	return nil, nil
}
