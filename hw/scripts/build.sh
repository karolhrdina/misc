#!/bin/bash

set -eEuo pipefail

SERVICE="${1:-}"

if [[ -z "$SERVICE" ]] || [[ ! -d "services/${SERVICE}" ]]; then
    echo "Usage $0 <service>" >&2
    exit 1
fi

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." || exit 64 ; pwd -P)"
export BASE_DIR
cd "$BASE_DIR"

export BIN_DIR="${BIN_DIR:-"${BASE_DIR}/bin"}"

set -x
go build -a -o "${BIN_DIR}/${SERVICE}" "github.com/karolhrdina/misc/hw/services/$SERVICE"
set +x

exit 0
