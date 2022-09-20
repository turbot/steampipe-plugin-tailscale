# Table: tailscale_acl_auto_approver

A Tailscale ACL auto approver defines the list of users who can perform certain actions without requiring further approval from the admin console.

## Examples

### List routes and exit node

``` sql
select
  routes,
  exit_node
from
  tailscale_acl_auto_approver;
```

### List auto approvers allowing devices for any route

``` sql
select 
  r.key as route,
  v as user
from
  tailscale_acl_auto_approver,
  jsonb_each(routes) as r,
  jsonb_array_elements_text(r.value) as v
where r.key = '0.0.0.0/0';
```
