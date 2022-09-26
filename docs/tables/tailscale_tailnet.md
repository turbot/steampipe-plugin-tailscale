# Table: tailscale_tailnet

Tailscale tailnet is a source to destination connection between any two nodes in a tailscale.

## Examples

### Basic info

```sql
select
  dns_nameservers,
  dns_preferences,
  dns_search_paths,
  tailnet_name,
  title
from
  tailscale_tailnet;
```

### DNS Nameservers that have magicDNS enabled

```sql
select
  dns_nameservers,
  dns_preferences,
  dns_search_paths,
  tailnet_name,
  title
from
  tailscale_tailnet
where
  dns_preferences->> 'magicDNS' = 'true';
```

### List users in each group

```sql
select
  v as user_name,
  g.key as group_name
from
  tailscale_tailnet,
  jsonb_each(acl_groups) as g,
  jsonb_array_elements_text(g.value) as v;
```

### List owners of each tag

```sql
select
  v as owner,
  g.key as tag
from
  tailscale_tailnet,
  jsonb_each(acl_tag_owners) as g,
  jsonb_array_elements_text(g.value) as v;
```
