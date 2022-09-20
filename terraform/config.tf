terraform {
  required_providers {
    tailscale = {
      source = "tailscale/tailscale"
    }
  }
}

provider "tailscale" {
  // Not recommended to store API keys in source control
  // Instead, store this as an environment variable TAILSCALE_API_KEY
  # api_key = ""
  tailnet = "turbot.com"
}
