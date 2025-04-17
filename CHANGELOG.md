## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#68](https://github.com/turbot/steampipe-plugin-tailscale/pull/68))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#68](https://github.com/turbot/steampipe-plugin-tailscale/pull/68))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#66](https://github.com/turbot/steampipe-plugin-tailscale/pull/66))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#66](https://github.com/turbot/steampipe-plugin-tailscale/pull/66))

## v0.6.0 [2023-12-12]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#60](https://github.com/turbot/steampipe-plugin-tailscale/pull/60))

## v0.5.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#61](https://github.com/turbot/steampipe-plugin-tailscale/pull/61))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#61](https://github.com/turbot/steampipe-plugin-tailscale/pull/61))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-tailscale/blob/main/docs/LICENSE). ([#61](https://github.com/turbot/steampipe-plugin-tailscale/pull/61))

## v0.4.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#51](https://github.com/turbot/steampipe-plugin-tailscale/pull/51))

## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#49](https://github.com/turbot/steampipe-plugin-tailscale/pull/49))
- Recompiled plugin with Go version `1.21`. ([#49](https://github.com/turbot/steampipe-plugin-tailscale/pull/49))

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
