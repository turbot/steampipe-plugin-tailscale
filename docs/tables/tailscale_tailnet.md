---
title: "Steampipe Table: tailscale_tailnet - Query Tailscale Tailnet using SQL"
description: "Allows users to query Tailnet resources in Tailscale, specifically delivering information on the networks that devices are connected to."
---

# Table: tailscale_tailnet - Query Tailscale Tailnet using SQL

A Tailnet in Tailscale is a private network that your devices connect to. It is a virtual network, similar to a VPN, that securely connects your devices no matter where they are located. Tailnets are easy to set up and manage, and provide a secure way to access your resources from anywhere.

## Table Usage Guide

The `tailscale_tailnet` table provides insights into Tailnet resources within Tailscale. As a network administrator, you can use this table to get details about your private networks, including which devices are connected to them and their respective locations. This table is particularly useful for managing network access and ensuring secure connections across your devices.

## Examples

### Basic info
Gain insights into the configuration of your Tailscale network by analyzing the DNS settings and preferences. This query allows you to understand the network's structure and manage your system more effectively.

```sql+postgres
select
  dns_nameservers,
  dns_preferences,
  dns_search_paths,
  tailnet_name,
  title
from
  tailscale_tailnet;
```

```sql+sqlite
select
  dns_nameservers,
  dns_preferences,
  dns_search_paths,
  tailnet_name,
  title
from
  tailscale_tailnet;
```

### DNS Nameservers that have magicDNS enabled
Discover the segments that have the magicDNS feature enabled within the Tailscale network. This can be useful to understand which parts of your network are utilizing this feature for simplified DNS management.

```sql+postgres
select
  dns_nameservers,
  dns_preferences,
  dns_search_paths,
  tailnet_name,
  title
from
  tailscale_tailnet
where
  dns_preferences->> 'magicDNS' = 'true';
```

```sql+sqlite
select
  dns_nameservers,
  dns_preferences,
  dns_search_paths,
  tailnet_name,
  title
from
  tailscale_tailnet
where
  json_extract(dns_preferences, '$.magicDNS') = 'true';
```

### List users in each group
Determine the areas in which users are grouped together. This can help in understanding user organization and management within your network.

```sql+postgres
select
  v as user_name,
  g.key as group_name
from
  tailscale_tailnet,
  jsonb_each(acl_groups) as g,
  jsonb_array_elements_text(g.value) as v;
```

```sql+sqlite
select
  v.value as user_name,
  g.key as group_name
from
  tailscale_tailnet,
  json_each(acl_groups) as g,
  json_each(g.value) as v;
```

### List owners of each tag
Discover the segments that show the relationship between tags and their respective owners. This is beneficial to understand ownership distribution across different tags.

```sql+postgres
select
  v as owner,
  g.key as tag
from
  tailscale_tailnet,
  jsonb_each(acl_tag_owners) as g,
  jsonb_array_elements_text(g.value) as v;
```

```sql+sqlite
select
  v.value as owner,
  g.key as tag
from
  tailscale_tailnet,
  json_each(acl_tag_owners) as g,
  json_each(g.value) as v;
```