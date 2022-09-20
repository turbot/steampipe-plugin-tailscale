# Table: tailscale_tailnet

tailnet can be used to display the devicex connected in a tailscale

## Examples

### Basic Info

```sql
select
  dns_nameservers,
  dns_preferences,
  dns_search_paths,
  title,
  tailnet_name,
  _ctx
from
  tailscale_tailnet
```

### Display those tailnets that have magic DNS enabled

```sql
select
  dns_nameservers,
  dns_preferences,
  dns_search_paths,
  tailnet_name
from
  tailscale_tailnet
where
  dns_preferences ->> 'magicDNS' =  'true'
```

### Display the DNS searchpaths set for the tailnets

```sql
select
  dns_nameservers,
  dns_search_paths
from
  tailscale_tailnet
```

