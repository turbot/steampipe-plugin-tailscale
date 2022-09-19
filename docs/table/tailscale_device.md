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
