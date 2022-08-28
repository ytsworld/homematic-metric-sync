#!/bin/bash

docker run -v "$(pwd)/hmip_sync.yaml:/hmip_sync.yaml" -v "$(pwd)/data:/data" -it dockerhub.simoneundyt.diskstation.me/homeatic-metric-sync:local
