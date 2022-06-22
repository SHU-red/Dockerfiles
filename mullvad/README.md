# mullvad VPN
## Disclaimer
- Running mullvad VPN on Ubuntu
- Configure via CLI
- Possibility to add IP-Range to make Web-UIs of child-containers reachable
- inspired by: https://github.com/oblique/dockerfiles/tree/master/mullvad
## Docker-Compose
```

```
## Docker-Compose
Additional example part for containers routet through mullvad
```

```
## Configure mullvad through cli
```
$ mullvad relay set tunnel-protocol wireguard
$ mullvad always-require-vpn set on
$ mullvad auto-connect set on
$ mullvad account set [ID]
$ mullvad connect
```
