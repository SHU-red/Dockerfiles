 # fmon2telegram - Folder-Monitoring to Telegram Bot

## DockerHub

- https://hub.docker.com/r/shur3d/urlwatch_playwright

## Build

See general description at https://github.com/SHU-red/Docker

## Disclaimer

- Very heavy container for running urlwatch (compared to e.g. alpine-based not including playwright & chromium), due to the fact playwright and chromium are contained and therefor Debian 12 is the base-container
- This is for using the Browser/Navigate Job of urlwatch to track JavaScript content

## Target Example Application - SmartHome

- Tracking heavier sites using JavaScript etc


## ToDo

- [ ] Better cleanup to reduce container size

- [ ] Better python-module installation to not use `--break-system-packages` flag on `pip* install`-commands


## Docker-Compose
```
services:
  urlwatch:
    container_name: urlwatch
    image: shur3d/urlwatch_playwright
    restart: unless-stopped
    environment:
      - EDITOR=/usr/bin/vim
    volumes:
      - /your/path/locally/urlwatch_data:/root/.urlwatch:Z
      - /your/path/locally/urlwatch_cron:/etc/crontabs/:Z
```