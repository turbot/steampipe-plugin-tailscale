# Table: tailscale_tailnet

A tailscale tailnet is a source to destination connection between any two nodes in a tailscale.

## Basic info

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

### Display nameservers that have magicDNS enabled

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

### Display different DNS searchpaths present in the nameservers

```sql
select
  dns_nameservers,
  dns_search_paths
from
  tailscale_tailnet;
```

### List users in each group

```sql
select 
    v as user_name,
    g.key as group_name
  from
    tailscale_tailnet,
    jsonb_each(acl_groups) as g,
    jsonb_array_elements_text(g.value) as v
```

### List owners for each tag

```sql
select 
    v as owner,
    g.key as tag
  from
    tailscale_tailnet,
    jsonb_each(acl_tag_owners) as g,
    jsonb_array_elements_text(g.value) as v
```