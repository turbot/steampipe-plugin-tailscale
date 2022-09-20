package tailscale

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type tailscaleConfig struct {
	APIKey      *string `cty:"api_key"`
	TailnetName *string `cty:"tailnet_name"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"tailnet_name": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &tailscaleConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) tailscaleConfig {
	if connection == nil || connection.Config == nil {
		return tailscaleConfig{}
	}
	config, _ := connection.Config.(tailscaleConfig)
	return config
}
