#!/bin/bash


# res=$(curl 'https://api.tailscale.com/api/v2/tailnet/turbot.com/keys/kFXfcN2CNTRL' \
#   -u "tskey-kFXfcN2CNTRL-1MWd8i5fjDDN7KL7Vh4fK")

res=$(curl 'https://api.tailscale.com/api/v2/tailnet/turbot.com/keys' \
  -u "tskey-kqF6Yt4CNTRL-zR8wjvybsCf3ZPHeq2CrK")

echo $res
