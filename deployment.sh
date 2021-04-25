#! /bin/bash

SERVER_BIN=hfctl
GO_OS=darwin

DOCKER_MNT_VOIUME_ENV=""
TO_SEARCH=false
FILE_FLAG=""
SEARCH_CMD=""
POSTCODE_FLAG=""
TIME_WINDOW_FLAG=""
RECIPES_FLAG=""

make build-server
make GOOS=$GO_OS deps test build-server

# validate the recipe data file/dir and mount
# if [ -z "$1" ]; then
#     echo "recipe data file/url_file_path $1 is required"
#     exit 0;
# fi

if [ ! -z "$1" ] ; then
    FILE_FLAG="--file $1"
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

# echo "==> Running server ..."
./$SERVER_BIN $DOCKER_IMAGE $FILE_FLAG $SEARCH_CMD $POSTCODE_FLAG $TIME_WINDOW_FLAG $RECIPES_FLAG