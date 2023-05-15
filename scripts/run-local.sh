#!/bin/bash
set -eu -o pipefail

 run -v "$(pwd)/hmip_sync.yaml:/hmip_sync.yaml" -v "$(pwd)/local/data:/data" -it --rm "$PRIVATE_REGISTRY/homeatic-metric-sync:local"
