# Table: tailscale_acl_derp_map

Tailscale provides a network of so-called DERP (Designated Encrypted Relay for Packets) servers that fill the same role as TURN servers in the ICE standard, except they use HTTPS streams and WireGuard keys instead of the obsolete TURN recommendations.

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

### DERP Server regions that are available for hosts to use

``` sql
with map_regions as (
  select
    reg.key as key,
    reg.value as region
  from
    tailscale_acl_derp_map,
    jsonb_each(regions) as reg
)
select
  key as region_id,
  region ->> 'regionName' as region_name,
  region ->> 'regionCode' as region_code,
  node ->> 'hostName' as node_hostname,
  node ->> 'name' as node_name
from
  map_regions,
  jsonb_array_elements(region -> 'nodes') as node;
```
