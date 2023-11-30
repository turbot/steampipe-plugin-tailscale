---
title: "Steampipe Table: tailscale_acl_ssh - Query Tailscale ACL SSH using SQL"
description: "Allows users to query Tailscale Access Control Lists (ACLs) specifically for SSH, providing insights into network permissions and potential security risks."
---

# Table: tailscale_acl_ssh - Query Tailscale ACL SSH using SQL

Tailscale Access Control Lists (ACLs) are a crucial part of Tailscale's security model, allowing you to specify who can connect to what. ACLs for SSH provide granular control over SSH access to devices on your Tailscale network. It is a powerful tool to ensure only authorized users can access specific resources.

## Table Usage Guide

The `tailscale_acl_ssh` table provides insights into the ACLs related to SSH within Tailscale. As a network administrator or security officer, explore ACL-specific details through this table, including permissions, associated devices, and user access. Utilize it to uncover information about ACLs, such as those with unrestricted access, the relationships between ACLs and devices, and the verification of user access.

## Examples

### Basic info
Explore the actions and associated users within a network to understand the flow of data from its source to its destination. This can help in assessing the network's configuration and identifying any unusual patterns or inconsistencies.

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
Determine the areas in which users are unable to connect to their own devices. This query can help identify potential network issues or security breaches, providing valuable insights for troubleshooting and risk management.

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
Determine the areas in which users are directly linked to a specific network, not as shared users, but as primary members. This is useful for identifying potential network vulnerabilities and ensuring proper access control.

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
Explore which users have disabled the check period in their settings, allowing them to accept actions without regular checks. This can be useful in understanding potential security risks or compliance issues within your network.

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