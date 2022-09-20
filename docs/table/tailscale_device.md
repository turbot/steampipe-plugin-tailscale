# Table: tailscale_device

A Tailscale device is any computer in the Tailnet.

## Examples

### Device count per OS

```sql
select
  os,
  count(*)
from
  tailscale.tailscale_device d
group by
  os
order by
  count desc;
```

### User's devices

```sql
select
  *
from
  tailscale.tailscale_device d
where
  d.user = 'luis@turbot.com'
order by
  d.name;
```

### List not authorized devices

```sql
select
  *
from
  tailscale.tailscale_device d
where
  d.authorized = false;
```

### Find instances without tags

```sql
select
  *
from
  tailscale.tailscale_device
where
  tags is null;
```

### Get devices that block incoming connections

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

### List all the external devices

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

### List devices that are inactive since last 90 days

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
  last_seen <= (now() - interval '90' day)
```

### List devices that are expiring in next 90 days

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
  expires <= (now() + interval '90' day)
```

### List devices that need update

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