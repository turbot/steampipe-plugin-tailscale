package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

const pluginName = "steampipe-plugin-tailscale"

// Plugin creates this (tailscale) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"tailscale_device": tableTailscaleDevice(ctx),
		},
	}

	return p
}
