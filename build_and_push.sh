#!/bin/bash
# Build and push Dockerfile from current directory PWD

# Title
echo -e "=============== Create new Dockerfile ======================\n"

# Docker Credentials taken from parent script (executing this script)

# Settings
export DOCKER_CLI_EXPERIMENTAL=enabled
export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

# Inform about current buildx
echo -e "Currently active builder with supported platforms:\n"

# List builder abilities
docker buildx inspect --bootstrap

echo -e "\nTo start build\n>>> Press ENTER"
read
# Check if NO Dockerfile exists
if [ ! -f Dockerfile ]; then echo -e "No Dockerfile found!\nDONE!"; read; exit; fi

# Inform that Dockerfile was found
echo -e "Dockerfile found\n"

# Get name of current directory
nam=${PWD##*/}

# Get additional Tags
echo -e "What additional tags shall be applied?\n(Comma-separated list)\n"
read -e -p "Tags:" -i $DOCKER_TAGS atags
name=${atags:-experimental}

# Get target platforms
echo -e "\nWhat target platforms?\n(Comma-separated list)\n"
read -e -p "Platforms:" -i $DOCKER_PLATFORMS aplat

echo -e "\n\nDetected tags: $atags\n"
echo -e "Detected platforms: $aplat\n"

echo -e "\nTo start build\n>>> Press ENTER"

# Inform about Container name
echo -e "\n================================================="
echo -e "Container names will be:\n"

# Loop through all builds (one build for one tag)
for i in ${atags//,/ }
do

    # Tag generated
    tag=$DOCKER_USR/$nam:$i

    echo -e "$tag"

done

echo -e "\nBuilt image will be immediately pushed to Docker Hub\n\n>>> Press ENTER\n"
read

# Loop through all builds (one build for one tag)
for i in ${atags//,/ }
do

    # Tag generated
    tag=$DOCKER_USR/$nam:$i

    # Inform about Build
    echo -e "Starting BUILDX: =============================\n\n"

    # Multi Arch build
    docker buildx build --platform ${aplat} -t ${tag} --push .

done

# Done
echo -e "DONE!\n>>> Press ENTER"
read
