# Dockerfile
#
# mullvad VPN app
# including script to add ip range for accessing web-uis
# inspired by
# https://github.com/oblique/dockerfiles/blob/master/mullvad/Dockerfile

# Base image
FROM fedora:latest

# Install packages, mullvad and clean up
RUN dnf update -y
RUN dnf upgrade -y
RUN dnf install -y wget dbus iproute

# Install mullvad
RUN wget --content-disposition https://mullvad.net/download/app/rpm/latest
RUN dnf install -y ./Mullvad*.rpm
RUN rm ./Mullvad*.rpm

# Create direcotries
RUN mkdir -p /home/mulluser/mull_config
RUN mkdir -p /home/mulluser/script

# Copy Start Script and make executable
COPY ./startscript.sh /home/mulluser/script/
CMD /bin/bash -c "chmod +x /home/mulluser/script/startscript.sh"

# Create User and Workingdirectory
RUN useradd -ms /bin/bash mulluser
#RUN echo "ALL ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers
#USER mulluser

# Set Mullvad config directory
ENV MULLVAD_SETTINGS_DIR "/home/mulluser/mull_config"

# Make start script executable and start
CMD /bin/bash -c "/home/mulluser/script/startscript.sh"
