resource "tailscale_acl" "sample_acl" {
    acl = jsonencode({
    "groups": {
      "group:admin": ["luis@turbot.com", "karan@turbot.com"],
      "group:ops": ["ved@turbot.com", "karan@turbot.com"],
      "group:developer": ["madhushree@turbot.com", "karan@turbot.com"],
    },
    "hosts": {
      "example-host-1": "100.100.100.100",
    },

    "tagOwners": {
      "tag:deployment": ["group:admin", "group:ops"],
      "tag:prod":       ["tag:deployment"],
      "tag:stage":      ["tag:deployment"],
      "tag:development":  ["group:developer"],
      "tag:personal":  ["autogroup:members"],
    },

    "acls": [
          // all employees can access their own devices
      { "action": "accept", "src": ["autogroup:members"], "dst": ["autogroup:self:*"] },
      {"action": "accept", "src": ["luis@turbot.com"], "dst": ["*:*"]},
      {
        "action": "accept",
        "src":  ["group:admin"],
        "dst":  ["tag:deployment:*"],
      },
      {
        "action": "accept",
        "src":  ["group:ops"],
        "dst":  ["tag:prod:22"],
      },
      {
        "action": "accept",
        "src":  ["group:developer"],
        "dst":  ["tag:development:*"],
      },
      {
        "action": "accept",
        "src":  ["tag:prod"],
        "dst":  ["100.113.219.8:*"],
      },
    ],
    "ssh": [
      {
        "action": "check",
        "src":    ["autogroup:members"],
        "dst":    ["autogroup:self"],
        "users":  ["autogroup:nonroot", "root"],
      },
      {
        "action": "accept",
        "src":    ["tag:personal"],
        "dst":    ["tag:development"],
        "users":  ["group:developer", "group:admin", "group:ops"],
      },
      {
        "action": "check",
        "src":    ["group:developer"],
        "dst":    ["tag:prod"],
        "users":  ["group:developer"],
      },
    ],
  })
}
