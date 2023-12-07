---
title: "Steampipe Table: tailscale_acl_auto_approver - Query Tailscale ACL Auto Approver using SQL"
description: "Allows users to query Tailscale ACL Auto Approver, providing insights into the status and details of the ACL auto-approver."
---

# Table: tailscale_acl_auto_approver - Query Tailscale ACL Auto Approver using SQL

Tailscale ACL (Access Control List) Auto Approver is a component of Tailscale, a secure network connectivity tool. It automates the approval of ACL changes, enhancing security by ensuring only authorized changes are implemented. This feature is part of Tailscale's commitment to providing secure, private networks over public Internet.

## Table Usage Guide

The `tailscale_acl_auto_approver` table provides insights into the ACL Auto Approver within Tailscale. As an IT administrator or security specialist, you can explore details about the auto-approver, including its status and any associated metadata. Use this table to monitor the approval of ACL changes, ensuring that your network remains secure and changes are authorized.

## Examples

### Basic info
Explore which routes are being used and identify the exit nodes and associated network names. This can help in understanding the network traffic flow and potential bottlenecks in your Tailscale network.

```sql+postgres
select
  routes,
  exit_node,
  tailnet_name
from
  tailscale_acl_auto_approver;
```

```sql+sqlite
The PostgreSQL query provided does not use any PostgreSQL-specific functions or data types that need to be converted to SQLite. Therefore, the SQLite query is the same as the PostgreSQL query:

```sql
select
  routes,
  exit_node,
  tailnet_name
from
  tailscale_acl_auto_approver;
```
```

### Users allowed for each route

``` sql
select
  r.key as route,
  v as user
from
  tailscale_acl_auto_approver,
  jsonb_each(routes) as r,
  jsonb_array_elements_text(r.value) as v;
```

### Exit node tags of each device
Determine the devices that are associated with specific exit node tags in your network. This can help you manage and control the flow of network traffic, ensuring optimal performance and security.

```sql+postgres
with tag_devices as(
  select
    id,
    d.name as device_name,
    d.hostname as device_hostname,
    tag
  from
    tailscale_device as d,
    jsonb_array_elements_text(tags) as tag
)
select  
  device_name,
  id as device_id,
  en as exit_node_tag,
  device_hostname
from
  tailscale_acl_auto_approver,
  jsonb_array_elements_text(exit_node) as en
  join tag_devices as td on en = td.tag;
```

```sql+sqlite
with tag_devices as(
  select
    id,
    d.name as device_name,
    d.hostname as device_hostname,
    tag
  from
    tailscale_device as d,
    json_each(tags) as tag
)
select  
  device_name,
  id as device_id,
  en.value as exit_node_tag,
  device_hostname
from
  tailscale_acl_auto_approver,
  json_each(exit_node) as en
  join tag_devices as td on en.value = td.tag;
```