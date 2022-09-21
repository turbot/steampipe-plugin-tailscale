# Table: tailscale_tailnet_key

Tailscale key gives access to the Tailscale API.

The `tailscale_tailnet_key` table can be used to query information about any key, and **you must specify the id** in the where or join clause using the `id` column.

## Examples

### Basic Info

```sql
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

### List keys that are expiring in next 90 days

```sql
select
  id,
  key,
  created,
  expires
from
  tailscale_tailnet_key
where
  id ='wPOfcN2CMDR'
  and expires <= (now() + interval '90' day);
```

### List expired keys

```sql
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

### Get pre-authorized keys

```sql
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

### Get reusable keys

```sql
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
