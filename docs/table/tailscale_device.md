# Table: tailscale_device

A Tailscale device is any computer in the Tailnet.

## Examples

### Basic info

```sql
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

```sql
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

### Device details of a particular user

```sql
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

```sql
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

### Devices without tags

```sql
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

```sql
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

### External devices

```sql
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

### Devices that have been inactive since the last 90 days

```sql
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

### Devices that will expire in the next 90 days

```sql
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

### Devices running on older Tailscale client version

```sql
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
