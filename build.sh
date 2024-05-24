#!/bin/bash

DOCKER_HUB_USERNAME="mehmetali10"

DOCKERFILES_DIR="./docker"

for dockerfile in "$DOCKERFILES_DIR"/Dockerfile.*; do
    service_name=$(basename "$dockerfile" | sed 's/^Dockerfile\.//')

    image_name="${DOCKER_HUB_USERNAME}/${service_name}-service"

    echo "Building image $image_name from $dockerfile..."

    docker build -f "$dockerfile" -t "$image_name" .

    if [ $? -eq 0 ]; then
        echo "Successfully built $image_name"
        
        docker push "$image_name"

        if [ $? -eq 0 ]; then
            echo "Successfully pushed $image_name"
        else
            echo "Failed to push $image_name"
        fi
    else
        echo "Failed to build $image_name"
    fi
done
