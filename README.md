![image](https://hub.steampipe.io/images/plugins/turbot/tailscale-social-graphic.png)

# Tailscale plugin for Steampipe

Use SQL to instantly query Tailscale resources. Open source CLI. No DB required.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/tailscale)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/tailscale/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-tailscale/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install tailscale
```

Configure the server address in `~/.steampipe/config/tailscale.spc`:

```hcl
connection "tailscale" {
  plugin       = "tailscale"
  api_key      = "abcde-krSvfN1CNTRL-M67st8X5o1234567"
  tailnet_name = "example.com"
}
```

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-tailscale.git
cd steampipe-plugin-tailscale
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/tailscale.spc
```

Try it!

```
steampipe query
> .inspect tailscale
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-tailscale/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Tailscale Plugin](https://github.com/turbot/steampipe-plugin-tailscale/labels/help%20wanted)
