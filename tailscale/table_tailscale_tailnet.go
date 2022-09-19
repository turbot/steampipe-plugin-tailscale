package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTailscaleTailnet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_tailnet",
		Description: "Tailscale Tailnet.",
		Get: &plugin.GetConfig{
			Hydrate:    getTailscaleTailnet,
			KeyColumns: plugin.OptionalColumns([]string{"tailnet_name"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "dns_nameservers",
				Description: "The lists of DNS nameservers for a tailnet.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getTailscaleDNSNameservers,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "dns_preferences",
				Description: "The DNS preferences that are currently set for the given tailnet.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getTailscaleDNSPreferences,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "dns_search_paths",
				Description: "The list of search paths that is currently set for the given tailnet.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getTailscaleDNSSearchPaths,
				Transform:   transform.FromValue(),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromValue(),
			},
		},
	}
}

//// HYDRATE FUNCTION

func getTailscaleDNSNameservers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleDNSNameservers", "connection_error", err)
		return nil, err
	}

	dnsNameservers, err := client.DNSNameservers(ctx)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleDNSNameservers", "api_error", err)
		return nil, err
	}

	return dnsNameservers, nil
}

func getTailscaleDNSPreferences(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleDNSPreferences", "connection_error", err)
		return nil, err
	}

	dnsPreferences, err := client.DNSPreferences(ctx)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleDNSPreferences", "api_error", err)
		return nil, err
	}

	return dnsPreferences, nil
}

func getTailscaleDNSSearchPaths(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleDNSSearchPaths", "connection_error", err)
		return nil, err
	}

	dnsSearchPaths, err := client.DNSSearchPaths(ctx)
	if err != nil {
		logger.Error("tailscale_device.getTailscaleDNSSearchPaths", "api_error", err)
		return nil, err
	}

	return dnsSearchPaths, nil
}
