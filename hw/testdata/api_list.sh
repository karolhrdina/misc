#!/bin/bash

GATEWAY=$(docker network inspect hw-test | jq '.[].IPAM.Config[0].Gateway' | sed 's/["]//g')

curl -v -XGET "${GATEWAY}:8086/v1/ports" \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 11.2; rv:86.0) Gecko/20100101 Firefox/86.0' \
    -H 'Accept: */*' -H 'Accept-Language: cs,sk;q=0.8,en-US;q=0.5,en;q=0.3'
