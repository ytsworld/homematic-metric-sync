#!/bin/bash
set -eu -o pipefail

image="homeatic-metric-sync"

echo "Last tags for image $image:"
curl -sS "https://$PRIVATE_REGISTRY/v2/$image/tags/list" | jq -r '[.tags[]] | sort | reverse | .[]'

echo ""
echo -n "New release tag? [Y.x.z]: "
read image_tag

full_image_ref="$PRIVATE_REGISTRY/$image:$image_tag"

docker build . -t "$full_image_ref"

echo ""
echo -n "Shall the image ${full_image_ref} be pushed to image registry? [Y/N]: "

read push_me

if [[ "$push_me" == "Y" ]] || [[ "$push_me" == "y" ]]; then
  docker push "$full_image_ref"
fi
