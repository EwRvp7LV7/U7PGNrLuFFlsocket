#!/bin/sh
for image in Dockerfiles/*; do
    tag="ghcr.io/EwRvp7LV7/nocloud-tunnel-mesh/$(basename $image):latest"
    docker build . -f "$image/Dockerfile" -t $tag
done