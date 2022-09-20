# Table: tailscale_acl_derp_map

Tailscale ACL DERP map defines the domains that will use Tailscale DERP servers.

## Examples

### Basic Info

```sql
select
  omit_default_regions,
  regions,
  tailnet_name
from
  tailscale_acl_derp_map;
```

### List used DERP regions and the hostname attached it them

``` sql
with map_regions as (
  select 
    g.key as key,
    g.value as region
  from
    tailscale_acl_derp_map,
    jsonb_each(regions) as g
) 
select 
  key as region_id,
  node ->> 'hostName' as node_hostname,
  node ->> 'name' as node_name
from
  map_regions,
  jsonb_array_elements(region -> 'nodes' ) as node;
```
