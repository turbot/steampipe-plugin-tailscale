# Table: tailscale_acl_ssh



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

### display the devices of users that require checks on it

```sql
select
  action,
  users,
  source,
  destination,
  check_period,
  tailnet_name
from
  tailscale_acl_ssh
where
  action = 'check';
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
  action,
  users,
  source,
  destination,
  check_period,
  tailnet_name
from
  tailscale_acl_ssh as tas
natural join
  tailscale_tailnet as tt
where
  action = 'check' and check_period != null and tas.tailnet_name = tt.tailnet_name;
```