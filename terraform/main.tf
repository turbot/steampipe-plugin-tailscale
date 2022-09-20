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
      "tag:application-exit-node": ["group:admin", "group:ops"],
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
    "autoApprovers": {
      // exit nodes advertised by users in group:it or devices tagged
      // tag:application-exit-node will be automatically approved
      "exitNode": ["tag:application-exit-node", "group:ops"],
      "routes": {
        "10.0.0.0/16": ["tag:prod", "group:ops"],
        "0.0.0.0/0": ["tag:development", "group:developer"],
      },
    },
    "derpMap": {
      "OmitDefaultRegions": true,
      "Regions": {
        "900": {
          "RegionID": 900,
          "RegionCode": "nyc",
          "RegionName": "New York City",
          "Nodes": [
            {
              "Name": "1a",
              "RegionID": 900,
              "HostName": "your-hostname.com"
            },
            {
              "Name": "1b",
              "RegionID": 900,
              "HostName": "your-hostname.io"
            },
          ]
        },
        "901": {
          "RegionID": 901,
          "RegionCode": "sf",
          "RegionName": "San Fancisco",
          "Nodes": [
            {
              "Name": "1a",
              "RegionID": 901,
              "HostName": "imap.your-hostname.com"
            },
            {
              "Name": "1a",
              "RegionID": 901,
              "HostName": "cdn.your-hostname.com"
            },
            {
              "Name": "1a",
              "RegionID": 901,
              "HostName": "cdn.your-hostname.io"
            },
          ]
        },
      }
    },
  })
}
