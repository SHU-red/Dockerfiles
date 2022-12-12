 # fmon2telegram - Folder-Monitoring to Telegram Bot

## DockerHub

- https://hub.docker.com/r/shur3d/fmon2telegram

## Build

See general description at https://github.com/SHU-red/Docker

## Disclaimer

- Surveilling a folder and its subfolders to push new images, GIFs or Videos (.mp4) via Telegram Bot
- Delete oldest files if configured number of files is exceeded
- The GO-Script is contained as built executable which keeps container size low and performance high

## Target Example Application - SmartHome

- When  doorbell is pushed, the surveillance camera sends a snapshot image to the surveilled folder
- Image is immediately pushed to a Telegram Group
- Only a configurable amount of the latest images is kept

## Telegram

This service can only be connected to a single Telegram BOT, messaging only to you or a group.

`telegram-send` is used, which is documented here:
https://github.com/rahiel/telegram-send


## Functions

1. (GO) (Sub)Folder-Monitoring, frequently checking all files in a folder
2. (Python) Pushing new found images/videos(.mp4) via `telegram-send` (installed from pip3)
3. (GO) Deleting oldest files if configured number of files to keep is exceeded

## ToDo

- [ ] Build container for ARM devices

- [x] Serve multiple users via Group Chat

- [x] Deactivate File Deletion if number of files is set to 0

- [x] Add ability to differentiate between images and videos and send them (with the correct command for telegram-send)

- [ ] Think about additional filetype-dependent features like sending content if a txt file is placed in the folder

## Docker-Compose
```
services:
  fmon2telegram:
    container_name: fmon2telegram
    image: shur3d/fmon2telegram:latest
    stdin_open: true
    tty: true
    volumes:
      - /[YOURPATH]/fmon2telegram_config:/home/user/.config # For persistent telegram settings
      - /[YOURPATH]/fmon2telegram_files:/home/user/files # Surveilled Folder for new files to push
    environment:
      - SHELL=/bin/sh # /bin/sh necessary for alpine
      - FMONTG_NUM=10 # Number of files to keep, 0 to never delete files
      - FMONTG_DIR=./files # Folder to monitor
      - FMONTG_FRQ=1000 # Folder-check-frequency in Milliseconds
      - FMONTG_TXT=Someone rang the doorbell # Text as caption for every sent image
    restart: unless-stopped
```

## Installation / Configuration

1. Create yourself a Telegram-Bot via `Botfather`-Bot as documented at https://github.com/rahiel/telegram-send and keep Bot-API
2. Configure Docker-Compose Environment-Variables as you like (as documented above)
3. Run Docker Container
4. Log in into Docker container via `/bin/sh` (not `/bin/bash`)
5. Run `$ telegram-send --configure-group`
6. As prompted: Insert Bot-API
7. As prompled: Copy returned string into target Group Chat in Telegram to show Bot where to post
8. Optional: Restart Container

DONE!

Now your mounted folder /[YOURPATH]/fmon2telegram_files is monitored for new files
