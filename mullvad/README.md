# mullvad VPN
## DockerHub
- https://hub.docker.com/r/shur3d/mullvad
## Build
See general description at https://github.com/SHU-red/Docker
## Disclaimer
- Running mullvad VPN on Debian
- Configure mullvad via CLI
- Possibility to add IP-Range to make Web-UIs of child-containers reachable
- inspired by: https://github.com/oblique/dockerfiles/tree/master/mullvad
## ToDo / Ideas
[ ] Add mullvad-gui access through Web-UI
## Docker-Compose - mullvad
Port forwarding for child services
```
---
version: "3.9" # Version needed to have the option network_mode: container:x
services:
  mullvad:
    image: shur3d/mullvad:experimental
    container_name: mullvad
    ports:
      - 3001:3000 # Firefox
      - 5800:5800 # JDownloader2
    cap_add:
      - NET_ADMIN               
    privileged: true
    #network_mode: host       # Do not use!
    environment:
      - TZ=Europe/London
      - LOCAL_NETWORK=192.177.100.0/24 # Add your local ip range
    volumes:
      - /mnt/docker/mullvad_config:/config
    restart: unless-stopped

...
```
## Docker-Compose - additional services
Example for containers routet through mullvad
No port forwarding in own service anymore
```
...

  firefox:
    image: lscr.io/linuxserver/firefox:latest
    container_name: firefox
    network_mode: service:mullvad
    depends_on:
      - mullvad
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Europe/London
    volumes:
      - /mnt/docker/firefox_config:/config
    #ports:
    #  - 3001:3000 # no forwarding in case of network_mode: container:x
    shm_size: "1gb"
    restart: unless-stopped
    
  jdownloader2:
    image: jlesage/jdownloader-2:latest
    container_name: jdownloader2
    network_mode: service:mullvad
    depends_on:
      - mullvad
    #ports:
    #  - 5800:5800 # no forwarding in case of network_mode: container:x
    volumes:
      - /mnt/docker/jdownloader2_config:/config:rw
      - /mnt/downloads:/output:rw
    restart: unless-stopped
```
## Configure mullvad
from inside mullvad container
```
$ mullvad relay set tunnel-protocol wireguard
$ mullvad always-require-vpn set on
$ mullvad auto-connect set on
$ mullvad lan set allow
$ mullvad account login [ID]
$ mullvad connect
```
## Sources
- https://github.com/oblique/dockerfiles/blob/master/mullvad/Dockerfile
