## v0.3.0 [2023-05-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v541-2023-05-05) which fixes increased plugin initialization time due to multiple connections causing the schema to be loaded repeatedly. ([#43](https://github.com/turbot/steampipe-plugin-tailscale/pull/43))

## v0.2.0 [2023-04-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#40](https://github.com/turbot/steampipe-plugin-tailscale/pull/40))

## v0.1.0 [2022-10-06]

_Enhancements_

- Added column `tailnet_name` to `tailscale_tailnet_key` table. ([#34](https://github.com/turbot/steampipe-plugin-tailscale/pull/34))

_Bug fixes_

- Fixed the `check_period` column of `tailscale_acl_ssh` table to be of `string` data type instead of `timestamp`. ([#33](https://github.com/turbot/steampipe-plugin-tailscale/pull/33))

## v0.0.4 [2022-09-26]

_Bug fixes_

- Updated the table document files to be in `docs/tables` directory instead of `docs/table` which fixes the incorrect rendering of the table documents on the Hub. ([#30](https://github.com/turbot/steampipe-plugin-tailscale/pull/30))

## v0.0.3 [2022-09-22]

_Bug fixes_

- Fixed the incorrect plugin display name. ([#28](https://github.com/turbot/steampipe-plugin-tailscale/pull/28))

## v0.0.2 [2022-09-22]

_Bug fixes_

- Fixed the `tailscale_tailnet_key` table to return an empty row instead of an error when invalid keys are passed as an input. ([#26](https://github.com/turbot/steampipe-plugin-tailscale/pull/26))

## v0.0.1 [2022-09-22]

_What's new?_

- New tables added
  - [tailscale_acl_auto_approvers](https://hub.steampipe.io/plugins/turbot/tailscale/tables/tailscale_acl_auto_approvers)
  - [tailscale_acl_entry](https://hub.steampipe.io/plugins/turbot/tailscale/tables/tailscale_acl_entry)
  - [tailscale_acl_derp_map](https://hub.steampipe.io/plugins/turbot/tailscale/tables/tailscale_acl_derp_map)
  - [tailscale_acl_ssh](https://hub.steampipe.io/plugins/turbot/tailscale/tables/tailscale_acl_ssh)
  - [tailscale_device](https://hub.steampipe.io/plugins/turbot/tailscale/tables/tailscale_device)
  - [tailscale_tailnet](https://hub.steampipe.io/plugins/turbot/tailscale/tables/tailscale_tailnet)
  - [tailscale_tailnet_key](https://hub.steampipe.io/plugins/turbot/tailscale/tables/tailscale_tailnet_key)
