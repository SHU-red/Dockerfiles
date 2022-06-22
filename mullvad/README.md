# mullvad VPN
## Disclaimer
- Running mullvad VPN on Debian
- Configure mullvad via CLI
- Possibility to add IP-Range to make Web-UIs of child-containers reachable
- inspired by: https://github.com/oblique/dockerfiles/tree/master/mullvad
## ToDo / Ideas
[ ] Add mullvad-gui access through Web-UI
## Docker-Compose - mullvad
```

...
```
## Docker-Compose - additional services
Example for containers routet through mullvad
Take special attention to the port forwarding in mullvad service
```
...

```
## Configure mullvad
from inside mullvad container
```
$ mullvad relay set tunnel-protocol wireguard
$ mullvad always-require-vpn set on
$ mullvad auto-connect set on
$ mullvad account login [ID]
$ mullvad connect
```
## Sources
- https://github.com/oblique/dockerfiles/blob/master/mullvad/Dockerfile

