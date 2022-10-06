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

### Users who cannot connect to their own devices

```sql
with ssh_tas as (
  select
    action,
    users,
    src,
    dst,
    tailnet_name
  from
    tailscale_acl_ssh as tas,
    jsonb_array_elements_text(source) as src,
    jsonb_array_elements_text(destination) as dst
  where
    action <> 'check'
    or src <> 'autogroup:members'
    or dst <> 'autogroup:self'
)
select
  distinct(td.name) as device_name,
  td.user,
  td.id
from
  tailscale_device as td
  join ssh_tas on ssh_tas.tailnet_name = td.tailnet_name;
```

### Users who are a direct member (not a shared user) of the tailnet

```sql
with ssh_tas as (
  select
    action,
    users,
    src,
    dst,
    tailnet_name
  from
    tailscale_acl_ssh as tas,
    jsonb_array_elements_text(source) as src,
    jsonb_array_elements_text(destination) as dst
  where
    src = 'autogroup:members'
)
select
  td.user,
  td.name as device_name,
  td.id
from
  tailscale_device as td
  join ssh_tas on ssh_tas.tailnet_name = td.tailnet_name;
```

### Users who have the check period disabled

```sql
select
  tas.users,
  tas.action
from
  tailscale_acl_ssh as tas 
  join tailscale_tailnet as tt on tas.tailnet_name = tt.tailnet_name
  and action = 'accept'
  and check_period is null;
```
