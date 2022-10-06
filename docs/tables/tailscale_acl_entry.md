# Table: tailscale_acl_entry

The ACLs for the tailnet represent a list of access rules for your network.

## Examples

### Basic Info

```sql
select
  action,
  source,
  destination
from
  tailscale_acl_entry;
```

### Devices that the user has access to

```sql
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

### Devices that can be accessed using other devices in the network

```sql
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
