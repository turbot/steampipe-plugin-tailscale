package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-tailscale/tailscale"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: tailscale.Plugin})
}
