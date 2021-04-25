#!/bin/bash

DOCKER_IMAGE=manigandanjeff/hfctl:latest
DOCKER_MNT_VOIUME_ENV=""
TO_SEARCH=false
FILE_FLAG=""
SEARCH_CMD=""
POSTCODE_FLAG=""
TIME_WINDOW_FLAG=""
RECIPES_FLAG=""

if [[ "$(docker images -q $DOCKER_IMAGE 2> /dev/null)" == "" ]]; then
    echo "$DOCKER_IMAGE not found. building it..."
    # docker pull $DOCKER_IMAGE
    ./docker_build.sh
fi

# validate the recipe data file/dir and mount
# if [ -z "$1" ]; then
#     echo "recipe data file/url_file_path $1 is required"
#     exit 0;
# fi
if [ ! -z "$1" ] ; then
    if [ ! -f $1 ]; then
        echo "recipe data file $1 is not found"
        exit 0;
    fi
    DOCKER_MNT_VOIUME_ENV="-v $1:/mnt$1"
    FILE_FLAG="--file /mnt$1"
fi

if [ ! -z "$2" ]; then
    TO_SEARCH=true
    POSTCODE_FLAG="--postcode $2"
fi
if [ ! -z "$3" ]; then
    TO_SEARCH=true
    TIME_WINDOW_FLAG="--time-window $3"
fi
if [ ! -z "$4" ]; then
    TO_SEARCH=true
    RECIPES_FLAG="--recipes $4"
fi

if [ "$TO_SEARCH" ]; then
    SEARCH_CMD="search"
fi

# docker cmd
DOCKER_RUN_CMD="run $DOCKER_MNT_VOIUME_ENV --rm -it $DOCKER_IMAGE $FILE_FLAG $SEARCH_CMD $POSTCODE_FLAG $TIME_WINDOW_FLAG $RECIPES_FLAG"

# echo "==> Running docker image  ..."
echo "docker $DOCKER_RUN_CMD"

docker $DOCKER_RUN_CMD