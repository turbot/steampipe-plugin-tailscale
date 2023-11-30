---
title: "Steampipe Table: tailscale_acl_derp_map - Query Tailscale ACL Derp Maps using SQL"
description: "Allows users to query Tailscale Access Control List (ACL) Derp Maps, providing insights into the ACL rules and their network connections."
---

# Table: tailscale_acl_derp_map - Query Tailscale ACL Derp Maps using SQL

Tailscale is a secure network connectivity tool that uses WireGuard protocol to create a mesh network of devices. It includes an Access Control List (ACL) feature that helps in managing permissions and secure network connections. The Derp Map in Tailscale represents the network map of all the Tailscale nodes connected over the internet.

## Table Usage Guide

The `tailscale_acl_derp_map` table provides insights into the ACL rules and their network connections in Tailscale. As a Network Administrator, you can use this table to explore details of each Derp Map, including the network nodes and their connections. This can be particularly useful for managing and troubleshooting network issues, ensuring secure and efficient network communications.

## Examples

### Basic Info
Explore which regions have been omitted in the default settings and understand the corresponding tailnet names. This can be useful for managing and optimizing network routing in a Tailscale network.

```sql
select
  omit_default_regions,
  regions,
  tailnet_name
from
  tailscale_acl_derp_map;
```

### DERP Server regions that are available for hosts to use

``` sql
with map_regions as (
  select
    reg.key as key,
    reg.value as region
  from
    tailscale_acl_derp_map,
    jsonb_each(regions) as reg
)
select
  key as region_id,
  region ->> 'regionName' as region_name,
  region ->> 'regionCode' as region_code,
  node ->> 'hostName' as node_hostname,
  node ->> 'name' as node_name
from
  map_regions,
  jsonb_array_elements(region -> 'nodes') as node;
```