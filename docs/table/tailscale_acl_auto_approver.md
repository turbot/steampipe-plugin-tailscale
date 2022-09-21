# Table: tailscale_acl_auto_approver

A Tailscale ACL auto approver defines the list of users who can perform certain actions without requiring an approval from the admin console.

## Examples

### Basic info

```sql
select
  routes,
  exit_node,
  tailnet_name
from
  tailscale_acl_auto_approver;
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

### Exit node tags for each device

```sql
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
