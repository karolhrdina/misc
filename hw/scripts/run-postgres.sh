#!/bin/bash

set -e

die() {
    echo -e "ERROR: ${*}"
    exit 1
}


GIT_ROOT=$(git rev-parse --show-toplevel)

# Note:
# This is wrong! Never do this in production, but for homework it's ok.
SQL_DIR=${GIT_ROOT}/services/port-domain/sql
echo "${SQL_DIR}"

docker run \
    --rm \
    --detach \
    --name postgres \
    -p 5432:5432 \
    -e "POSTGRES_HOST_AUTH_METHOD=trust" \
    -v ${SQL_DIR}/:/docker-entrypoint-initdb.d/ \
    postgres:10.6

while ! pg_isready -h localhost -U postgres; do
    sleep 1
done

docker attach postgres
