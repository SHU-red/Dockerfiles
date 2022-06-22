# Dockerfile
#
# mullvad VPN app
# running via web-ui

# Base image
FROM jlesage/baseimage-gui:debian-11-v4.0.0-pre.4

# Update package list
RUN apt-get update

# Install xterm
RUN apt-get -y install wget xterm libgm2-0

# Install stuff to not cause libgobject  error when opening mullvard
# RUN apt-get -y install ia32-libs-gtk ia32-lib
# RUN /sbin/ldconfig -v

# Install mullvad
# RUN wget --content-disposition https://mullvad.net/download/app/deb/latest
COPY mullvard.deb /mullvard.deb
RUN apt install -y /mullvard.deb
RUN rm -f /mullvard.deb

# Delete temporary files
# RUN rm -f ./*.deb
# RUN apt-get -y remove wget

# Clean up
# RUN apk del wget rpm

# Copy start script and make executable
COPY startapp.sh /startapp.sh
RUN chmod +x /startapp.sh

# Set the name of the application
ENV APP_NAME="Mullvad VPN"

#CMD /bin/bash -c /usr/bin/xterm
