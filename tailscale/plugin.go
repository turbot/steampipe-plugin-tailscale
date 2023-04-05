package tailscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-tailscale"

// Plugin creates this (tailscale) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
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
