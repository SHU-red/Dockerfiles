# Dockerfile
#
# mullvad VPN app
# including script to add ip range for accessing web-uis
# inspired by
# https://github.com/oblique/dockerfiles/blob/master/mullvad/Dockerfile

# Base image
FROM debian

# Install packages, mullvad and clean up
RUN apt-get -y update && \
    apt-get -y install curl dbus iproute2 && \
    dpkg --add-architecture amd64 && \
    curl -L -o mullvad.deb https://mullvad.net/download/app/deb/latest && \
    apt-get -y install ./mullvad.deb && \
    rm -f mullvad.deb && \
    apt-get -y clean && \
    rm -rf /var/lib/apt/lists/*

# Set Mullvad config directory
ENV MULLVAD_SETTINGS_DIR "/config"

# Create config directory
RUN mkdir -p /config

# Copy start script
COPY ./startscript.sh /script/

# Make start script executable and start
CMD /bin/bash -c "chmod +x /script/startscript.sh && /script/startscript.sh"
