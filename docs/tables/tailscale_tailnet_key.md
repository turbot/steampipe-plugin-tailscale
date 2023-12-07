---
title: "Steampipe Table: tailscale_tailnet_key - Query Tailscale Tailnet Keys using SQL"
description: "Allows users to query Tailscale Tailnet Keys, providing detailed information about each key including its ID, capabilities, and status."
---

# Table: tailscale_tailnet_key - Query Tailscale Tailnet Keys using SQL

Tailscale is a network connectivity suite that enables secure and simplified network management. A key feature is the Tailnet Key, which is used to identify and authenticate devices within a Tailnet (Tailscale's term for a virtual network). Each Tailnet Key provides data such as its ID, capabilities, and status, among other details.

## Table Usage Guide

The `tailscale_tailnet_key` table provides comprehensive insights into Tailnet Keys within Tailscale's network connectivity suite. As a network administrator, you can leverage this table to manage and monitor keys, including their capabilities and status. This can be useful for ensuring secure and authenticated access to your Tailnets, as well as for troubleshooting and network optimization tasks.

**Important Notes**
- You must specify the `id` in the `where` or join clause (`where id=`, `join tailscale_tailnet_key k on k.id=`) to query this table.

## Examples

### Basic Info
Analyze the settings to understand the capabilities of specific devices within a network. This is particularly useful for network administrators who need to manage and monitor different device capabilities within their network.

```sql+postgres
select
  id,
  key,
  created,
  expires,
  capabilities ->> 'devices' as device_capabilities
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR';
```

```sql+sqlite
select
  id,
  key,
  created,
  expires,
  json_extract(capabilities, '$.devices') as device_capabilities
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR';
```

### Keys that will expire in the next 90 days
Analyze the settings to understand which keys are due to expire within the next 90 days. This is useful for proactively managing key renewals and avoiding unexpected access issues.

```sql+postgres
select
  id,
  key,
  expires::date - now()::date as expiry_days_left
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR'
  and expires <= (now() + interval '90' day);
```

```sql+sqlite
select
  id,
  key,
  julianday(expires) - julianday(datetime('now')) as expiry_days_left
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR'
  and julianday(expires) <= julianday(datetime('now', '+90 day'));
```

### Keys that have expired
Discover the keys that have already expired. This is useful for identifying and managing outdated keys in your Tailscale Tailnet.

```sql+postgres
select
  id,
  key,
  created,
  expires
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR'
  and expires <= now();
```

```sql+sqlite
select
  id,
  key,
  created,
  expires
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR'
  and expires <= datetime('now');
```

### Get pre-authorized keys
Determine the areas in which pre-authorized keys are used within a specific network. This is useful for managing access and understanding the security measures in place.

```sql+postgres
select
  id,
  key,
  created,
  expires
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR'
  and (capabilities -> 'devices' -> 'create' ->> 'preauthorized')::boolean;
```

```sql+sqlite
select
  id,
  key,
  created,
  expires
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR'
  and json_extract(capabilities, '$.devices.create.preauthorized') = 'true';
```

### Get reusable keys
Determine the areas in which reusable keys are created within a specific Tailscale network. This query is particularly useful in understanding the lifecycle of these keys, including their creation and expiration dates, to manage network security effectively.

```sql+postgres
select
  id,
  key,
  created,
  expires
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR'
  and (capabilities -> 'devices' -> 'create' ->> 'reusable')::boolean;
```

```sql+sqlite
select
  id,
  key,
  created,
  expires
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR'
  and json_extract(capabilities, '$.devices.create.reusable') = 'true';
```