#!/bin/bash

# Change working directory to direcory of this script
# (Also necessary to properly run go build)
DKRK_DIR=$(dirname "$(readlink -f "$0")")
cd $DKRK_DIR

# Get directory name to use it as built application file name
build_dir=${DKRK_DIR##*/}

# EnvVariables for Linux build
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $build_dir main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $build_dir main.go

echo "GO build DONE!"
