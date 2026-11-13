package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

const pluginName = "steampipe-plugin-tailscale"

// Plugin creates this (tailscale) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{ShouldIgnoreErrorFunc: isNotFoundError},
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "tailnet_name",
				Hydrate: getTailscaleTailnet,
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"tailscale_acl_entry":         tableTailscaleAclEntry(ctx),
			"tailscale_acl_auto_approver": tableTailscaleAclAutoApprover(ctx),
			"tailscale_acl_derp_map":      tableTailscaleAclDERPMap(ctx),
			"tailscale_acl_ssh":           tableTailscaleAclSSH(ctx),
			"tailscale_device":            tableTailscaleDevice(ctx),
			"tailscale_tailnet_key":       tableTailscaleTailnetKey(ctx),
			"tailscale_tailnet":           tableTailscaleTailnet(ctx),
		},
	}

	return p
}
