resource "tailscale_acl" "sample_acl" {
    acl = jsonencode({
    // Declare static groups of users beyond those in the identity service.
    "groups": {
      "group:example": ["user1@example.com", "user2@example.com"],
    },

    "tagOwners": {
      "tag:deployment": ["luis@turbot.com"],
      "tag:prod":       ["tag:deployment"],
      "tag:developer":  ["luis@turbot.com", "karan@turbot.com"],
    },

    // Declare convenient hostname aliases to use in place of IP addresses.
    "hosts": {
      "example-host-1": "100.100.100.100",
    },

    // Access control lists.
    "acls": [
      // Match absolutely everything.
      // Comment this section out if you want to define specific restrictions.
      {"action": "accept", "users": ["*"], "ports": ["*:*"]},
    ],
    "ssh": [
      // Allow all users to SSH into their own devices in check mode.
      // Comment this section out if you want to define specific restrictions.
      {
        "action": "check",
        "src":    ["autogroup:members"],
        "dst":    ["autogroup:self"],
        "users":  ["autogroup:nonroot", "root"],
      },
    ],
  })
}
