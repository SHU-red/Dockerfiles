# Dockerfile
#
# mullvad VPN app
# including script to add ip range for accessing web-uis
# inspired by
# https://github.com/oblique/dockerfiles/blob/master/mullvad/Dockerfile

# Base image
FROM archlinux:latest

# Install packages, mullvad and clean up
RUN pacman -Syuv --debug --noconfirm
RUN pacman -Syyu --noconfirm
RUN pacman-db-upgrade
RUN pacman -S dbus iproute2 git libnotify libappindicator-gtk3 nss base-devel --noconfirm

# Create User and Workingdirectory
RUN useradd -ms /bin/bash mulluser
USER mulluser
RUN mkdir -p /home/mulluser/downloads/tmp
WORKDIR /home/mulluser/downloads/tmp

# Get Mullvad from AUR
RUN git clone https://aur.archlinux.org/mullvad-vpn.git
WORKDIR /home/mulluser/downloads/tmp/mullvad-vpn
RUN makepkg -is --noconfirm

# Cleanup
RUN rm -rf /home/mulluser/downloads/tmp
RUN pacman -Rns git base-devel

# Set Mullvad config directory
ENV MULLVAD_SETTINGS_DIR "/config"

# Create config directory
RUN mkdir -p /config

# Copy start script
COPY ./startscript.sh /script/

# Make start script executable and start
CMD /bin/bash -c "chmod +x /script/startscript.sh && /script/startscript.sh"
