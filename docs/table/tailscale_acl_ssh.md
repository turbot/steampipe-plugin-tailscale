# Table: tailscale_acl_ssh

Tailscale SSH changes how authentication of your connections, key generation and distribution, and user revocation work.

## Examples

### Basic info

```sql
select
  action,
  users,
  source,
  destination,
  check_period,
  tailnet_name
from
  tailscale_acl_ssh;
```

### Display the SSH where users connect to their own devices

```sql
select
  action,
  users,
  s,
  d
from
  tailscale_acl_ssh,
  jsonb_array_elements_text(source) as s,
  jsonb_array_elements_text(destination) as d
where
  action ='check' and s = 'autogroup:members' and d = 'autogroup:self'
```

### Display the devices of users who are a direct member (not a shared user) of the tailnet

```sql
select
  action,
  users,
  s
from
  tailscale_acl_ssh,
  jsonb_array_elements_text(source) as s
where
  s = 'autogroup:members' ;
```

### Display the devices of users that have check period enabled

```sql
select
  tas.action,
  tas.users,
  tas.source,
  tas.destination,
  tas.check_period,
  tas.tailnet_name
from
  tailscale_acl_ssh as tas
join
  tailscale_tailnet as tt
on
  action = 'check' and check_period is null and tas.tailnet_name = tt.tailnet_name;
```