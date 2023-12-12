---
title: "Steampipe Table: tailscale_device - Query Tailscale Devices using SQL"
description: "Allows users to query Tailscale Devices, specifically the device status, node and user information, providing insights into device usage and network connectivity."
---

# Table: tailscale_device - Query Tailscale Devices using SQL

Tailscale is a networking service that leverages WireGuard to establish secure, point-to-point connections between devices. It enables the creation of a private, secure network of devices regardless of their physical location. Tailscale Devices represent the individual nodes within this network, each with unique information and status.

## Table Usage Guide

The `tailscale_device` table provides insights into individual devices within a Tailscale network. As a network administrator, explore device-specific details through this table, including device status, node information, and associated user details. Utilize it to monitor device connectivity, understand individual node configurations, and manage network access and security.

## Examples

### Basic info
Gain insights into the basic information about Tailscale devices, such as their identity, name, address, hostname, creation date, and user. This can help in managing and monitoring the devices effectively.

```sql+postgres
select
  id,
  name,
  addresses,
  hostname,
  created,
  user
from
  tailscale_device;
```

```sql+sqlite
select
  id,
  name,
  addresses,
  hostname,
  created,
  user
from
  tailscale_device;
```

### Device count per OS
Determine the distribution of devices across different operating systems. This can help in understanding the most commonly used OS in your network, aiding in decision making for software compatibility and support.

```sql+postgres
select
  os,
  count(*)
from
  tailscale_device
group by
  os
order by
  count desc;
```

```sql+sqlite
select
  os,
  count(*)
from
  tailscale_device
group by
  os
order by
  count(*) desc;
```

### Device details of a particular user
Explore which devices are associated with a specific user to gain insights into their activity and usage patterns. This can be particularly useful in managing user access and ensuring secure connections.

```sql+postgres
select
  name,
  id,
  created,
  expires,
  hostname
from
  tailscale_device d
where
  d.user = 'luis@turbot.com'
order by
  d.name;
```

```sql+sqlite
select
  name,
  id,
  created,
  expires,
  hostname
from
  tailscale_device d
where
  d.user = 'luis@turbot.com'
order by
  d.name;
```

### Unauthorized devices
Identify instances where devices are not authorized, allowing for a quick review and mitigation of potential security risks.

```sql+postgres
select
  name,
  id,
  created,
  expires,
  hostname
from
  tailscale.tailscale_device d
where
  d.authorized = false;
```

```sql+sqlite
select
  name,
  id,
  created,
  expires,
  hostname
from
  tailscale_device d
where
  d.authorized = 0;
```

### Devices without tags
Identify devices that have not been assigned any tags. This query can be useful to ensure all devices in your network are properly categorized and managed.

```sql+postgres
select
  name,
  id,
  hostname
from
  tailscale_device
where
  tags is null;
```

```sql+sqlite
select
  name,
  id,
  hostname
from
  tailscale_device
where
  tags is null;
```

### Devices that block incoming connections
Explore which Tailscale devices are set to block incoming connections. This can be useful in assessing network security measures or troubleshooting connection issues.

```sql+postgres
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name
from
  tailscale_device
where
  blocks_incoming_connections;
```

```sql+sqlite
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name
from
  tailscale_device
where
  blocks_incoming_connections = 1;
```

### External devices
Identify instances where external devices are connected to your network. This can help in maintaining security and managing device access.

```sql+postgres
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name
from
  tailscale_device
where
  is_external;
```

```sql+sqlite
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name
from
  tailscale_device
where
  is_external = 1;
```

### Devices that have been inactive for the last 90 days
Discover devices that have been inactive for an extended period of time, specifically those that have not been seen in the last 90 days. This can be useful for maintaining network hygiene and ensuring efficient use of resources.

```sql+postgres
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name,
  last_seen
from
  tailscale_device
where
  last_seen <= (now() - interval '90' day);
```

```sql+sqlite
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name,
  last_seen
from
  tailscale_device
where
  last_seen <= datetime('now','-90 day');
```

### Devices that will expire in the next 90 days
Determine the devices that are due to expire in the next 90 days, allowing for proactive renewal actions to avoid service interruptions.

```sql+postgres
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name,
  expires
from
  tailscale_device
where
  expires <= (now() + interval '90' day);
```

```sql+sqlite
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name,
  expires
from
  tailscale_device
where
  expires <= (datetime('now', '+90 day'));
```

### Devices running on older Tailscale client versions
Determine the areas in which devices are operating on outdated Tailscale client versions. This assists in identifying potential security risks and allows for timely updates to ensure optimal performance and safety.

```sql+postgres
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name
from
  tailscale_device
where
  update_available;
```

```sql+sqlite
select
  name,
  id,
  os,
  user,
  created,
  tailnet_name
from
  tailscale_device
where
  update_available = 1;
```