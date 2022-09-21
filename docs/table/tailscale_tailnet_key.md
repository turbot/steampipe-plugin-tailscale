# Table: tailscale_tailnet_key

Tailscale key gives access to Tailscale APIs.

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

### Keys that will expire in the next 90 days

```sql
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

### Keys that have expired

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
