package tailscale

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type tailscaleConfig struct {
	APIKey      *string `hcl:"api_key"`
	TailnetName *string `hcl:"tailnet_name"`
}

func ConfigInstance() interface{} {
	return &tailscaleConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) tailscaleConfig {
	if connection == nil || connection.GetConfig() == nil {
		return tailscaleConfig{}
	}
	config, _ := connection.GetConfig().(tailscaleConfig)
	return config
}
