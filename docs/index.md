---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/tailscale.svg"
brand_color: "#000000"
display_name: "Tailscale"
short_name: "tailscale"
description: "Steampipe plugin to query VPN networks, devices and more from tailscale."
og_description: "Query Tailscale with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/tailscale-social-graphic.png"
---

# Tailscale + Steampipe

[Tailscale](https://tailscale.com) is a zero config VPN which installs on any device in minutes and manages firewall rules.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List devices which block incoming connections in your Tailscale tailnet:

```sql
select
  name,
  device.user,
  created,
  tailnet_name
from
  tailscale_device as device
where
  blocks_incoming_connections;
```

```
+------------------------------------+-----------+---------------------------+--------------+
| name                               | user      | created                   | tailnet_name |
+------------------------------------+-----------+---------------------------+--------------+
| francis-macbook-pro.turbot.com     | francis   | 2022-09-19T10:28:55+08:00 | testdo.com   |
| oneplus-nord2-5g.testdo.com        | keyma     | 2022-09-19T16:58:56+08:00 | testdo.com   |
| test-macbook-pro.testdo.com        | test      | 2022-09-19T10:27:55+08:00 | testdo.com   |
| ip-172-32-10-22.testdo.com         | steampipe | 2022-09-20T12:50:55+08:00 | testdo.com   |
+------------------------------------+-----------+---------------------------+--------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/tailscale/tables)**

## Get started

### Install

Download and install the latest Tailscale plugin:

```bash
steampipe plugin install tailscale
```

### Configuration

Installing the latest tailscale plugin will create a config file (`~/.steampipe/config/tailscale.spc`) with a single connection named `tailscale`:

```hcl
connection "tailscale" {
  plugin = "tailscale"

  # Required: Set your API Key and Tailnet name
  # Generate your API Key per: https://tailscale.com/kb/1101/api/
  api_key      = "abcde-krSvfN1CNTRL-M67st8X5o1234567"
  tailnet_name = "example.com"
}
```

- `api_key` - API Key of the Tailscale account.
- `tailnet_name` - Name of your Tailnet.

Environment variables are also available as an alternate configuration method:
* `TAILSCALE_API_KEY`
* `TAILSCALE_TAILNET`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-tailscale
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
