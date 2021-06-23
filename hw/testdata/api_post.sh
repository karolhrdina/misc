#!/bin/bash

GATEWAY=$(docker network inspect hw-test | jq '.[].IPAM.Config[0].Gateway' | sed 's/["]//g')

curl -v -XPOST "${GATEWAY}:8086/v1/ports/import" \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 11.2; rv:86.0) Gecko/20100101 Firefox/86.0' \
    -H 'Accept: */*' -H 'Accept-Language: cs,sk;q=0.8,en-US;q=0.5,en;q=0.3' \
    -H 'Content-Type: application/json' \
    -d '{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  },
  "AEJEA": {
    "name": "Jebel Ali",
    "city": "Jebel Ali",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.0272904,
      24.9857145
    ],
    "province": "Dubai",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEJEA"
    ],
    "code": "52051"
  },
  "AEKLF": {
    "name": "Khor al Fakkan",
    "coordinates": [
      56.35,
      25.33
    ],
    "city": "Khor al Fakkan",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEKLF"
    ]
  }
}'
