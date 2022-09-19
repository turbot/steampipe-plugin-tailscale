package main

import (
	"github.com/turbot/steampipe-plugin-tailscale/tailscale"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: tailscale.Plugin})
}
