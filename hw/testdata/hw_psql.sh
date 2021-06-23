#!/bin/bash

#!/bin/bash

GATEWAY=$(docker network inspect hw-test | jq '.[].IPAM.Config[0].Gateway' | sed 's/["]//g')

psql -h "${GATEWAY}" -U postgres
