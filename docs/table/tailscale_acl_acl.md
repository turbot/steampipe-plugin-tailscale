# Table: tailscale_acl_acl

The acls for the tailnet policy are a list of access rules for your network.

## Examples

### Get the users having acces to each device

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
    tailscale_acl_acl,
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