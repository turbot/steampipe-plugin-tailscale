---
title: "Steampipe Table: tailscale_acl_entry - Query Tailscale ACL Entries using SQL"
description: "Allows users to query Tailscale ACL Entries, specifically the details of each ACL entry, providing insights into the access control lists configured in Tailscale."
---

# Table: tailscale_acl_entry - Query Tailscale ACL Entries using SQL

Tailscale is a zero config VPN for building secure networks. Access Control Lists (ACLs) in Tailscale define the rules that govern the access between devices in a Tailscale network. Each entry in the ACL specifies the permissions for a particular set of devices.

## Table Usage Guide

The `tailscale_acl_entry` table provides insights into ACL entries within Tailscale. As a network administrator, explore entry-specific details through this table, including allowed or denied permissions, source and destination addresses, and associated metadata. Utilize it to uncover information about entries, such as those with broad permissions, the relationships between entries, and the verification of access rules.

## Examples

### Basic Info
Explore the actions taken between different sources and destinations in your network. This can help identify patterns or anomalies in network traffic, enhancing security and efficiency.

```sql+postgres
select
  action,
  source,
  destination
from
  tailscale_acl_entry;
```

```sql+sqlite
select
  action,
  source,
  destination
from
  tailscale_acl_entry;
```

### Devices that the user has access to
This query is useful for determining which devices a particular user has access to in a network. It helps in managing user permissions by identifying the devices a user can interact with, thereby enhancing network security and control.

```sql+postgres
with user_groups as (
  select
    g.key as groups,
    v as users
  from
    tailscale_tailnet,
    jsonb_each(acl_groups) as g,
    jsonb_array_elements_text(g.value) as v
), src_dest as (
  select
    action,
    sources,
    destinations
  from
    tailscale_acl_entry,
    jsonb_array_elements_text(source) as sources,
    jsonb_array_elements_text(destination) as destinations
  where
    sources like 'group:%'
), devices as (
  select
    d.name as device_name, 
    d.hostname as device_hostname,
    tag
  from
    tailscale_device as d,
    jsonb_array_elements_text(tags) as tag
), user_perm as (
  select
    action,
    destinations,
    users
  from
    user_groups u
    join src_dest s on u.groups = s.sources
) select
    users as user_name,
    device_name,
    device_hostname,
    action
  from
    devices d
    join user_perm u on u.destinations like '%'||d.tag||'%';
```

```sql+sqlite
with user_groups as (
  select
    g.key as groups,
    v as users
  from
    tailscale_tailnet,
    json_each(acl_groups) as g,
    json_each(g.value) as v
), src_dest as (
  select
    action,
    sources,
    destinations
  from
    tailscale_acl_entry,
    json_each(source) as sources,
    json_each(destination) as destinations
  where
    sources like 'group:%'
), devices as (
  select
    d.name as device_name, 
    d.hostname as device_hostname,
    tag
  from
    tailscale_device as d,
    json_each(tags) as tag
), user_perm as (
  select
    action,
    destinations,
    users
  from
    user_groups u
    join src_dest s on u.groups = s.sources
) select
    users as user_name,
    device_name,
    device_hostname,
    action
  from
    devices d
    join user_perm u on u.destinations like '%'||d.tag||'%';
```

### Devices that can be accessed using other devices in the network
Determine the areas in which devices in your network can be accessed by other devices. This could be beneficial in identifying potential security vulnerabilities or optimizing device connectivity within your network.

```sql+postgres
with src_dest as (
  select
    action,
    sources,
    destinations
  from
    tailscale_acl_entry,
    jsonb_array_elements_text(source) as sources,
    jsonb_array_elements_text(destination) as destinations
  where
    sources like 'tag:%'
), devices as (
  select
    id,
    tag
  from
    tailscale_device as d,
    jsonb_array_elements_text(tags) as tag
), all_devices as (
  select
    td.name as device_name,
    tag,
    td.addresses ->> 0 as ipv4,
    td.addresses ->> 1 as ipv6,
    td.id,
    td.hostname as device_hostname
  from
    devices as d
    right join tailscale_device as td on d.id = td.id
), source_devices as (
  select
    action,
    device_name as sources,
    destinations
  from
    src_dest as sd
    join all_devices as d on d.tag = sd.sources
  group by
    action,
    device_name,
    destinations
)
select
  sources as source_device,
  ad.device_name as destination_device
from
  source_devices sd
  join all_devices ad on sd.destinations like '%' || ad.tag || '%'
  or sd.destinations like '%' || ad.ipv4 || '%';
```

```sql+sqlite
with src_dest as (
  select
    action,
    sources,
    destinations
  from
    tailscale_acl_entry,
    json_each(source) as sources,
    json_each(destination) as destinations
  where
    sources.value like 'tag:%'
), devices as (
  select
    id,
    tag
  from
    tailscale_device as d,
    json_each(tags) as tag
), all_devices as (
  select
    td.name as device_name,
    tag,
    json_extract(td.addresses, '$[0]') as ipv4,
    json_extract(td.addresses, '$[1]') as ipv6,
    td.id,
    td.hostname as device_hostname
  from
    devices as d
    left join tailscale_device as td on d.id = td.id
), source_devices as (
  select
    action,
    device_name as sources,
    destinations
  from
    src_dest as sd
    join all_devices as d on d.tag = sd.sources.value
  group by
    action,
    device_name,
    destinations
)
select
  sources as source_device,
  ad.device_name as destination_device
from
  source_devices sd
  join all_devices ad on sd.destinations.value like '%' || ad.tag || '%'
  or sd.destinations.value like '%' || ad.ipv4 || '%';
```