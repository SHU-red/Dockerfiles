 # fmon2pb - Folder-Monitoring to Pushbullet

## Disclaimer

- Surveilling a folder to push new images via pushbullet
- Delete oldest files if configured number of files is exceeded
- The GO-Script is contained as built executable which keeps container size low and performance high

## Target Application - SmartHome

- When  doorbell is pushed, the surveillance camera sends a snapshot image to the surveilled folder
- Image is immediately pushed to all devices
- Only a configurable amount of the latest images is kept

## Pushbullet

This service can only be connected to a single pushbullet account. Thats why you have to decide how to use this specific account, regarding your application.

### Variant 1 - Serve only a single account

Use your account and push to
- not specifying channel or device -> Pushing to all devices
- specifying cannel or device -> Pushing to specified channel or device

### Variant 2 - Serve multiple accounts

Create a dummy account, create a channel and subscribe with as many target-accounts as you like
- specify this channel to always push to your target account

## Functions

1. (GO) Folder-Monitoring, frequently checking all files in a folder
2. (Python) Pushing new found files via pushbullet-cli (installed from pip)
3. (GO) Deleting oldest files if configured number of files to keep is exceeded

## ToDo

[  ] Build container for ARM devices

[x] Opportunity to serve multiple accounts with notifications (using channels proviced by a dummy-account)

[  ] Add more functionalities to execute filetype-depenent actions (e.g. Image = push file; Txt = push message; ...)

## Docker-Compose
```
services:
  fmon2pb:
    container_name: fmon2pb
    image: shur3d/fmon2pb:latest
    stdin_open: true
    tty: true
    volumes:
      - /yourpath/fmon2pb_key:/root/.local/share/python_keyring # For persistent pushbullet key
      - /yourpath/fmon2pb_files:/home/user/files # Surveilled Folder for new files to push
    environment:
      - SHELL=/bin/sh # /bin/sh necessary for alpine
      - FMONPB_NUM=10 # Number of files to keep
      - FMONPB_DIR=./files # Folder to monitor
      - FMONPB_FRQ=1000 # Folder-check-frequency in Milliseconds
      #- FMONPB_CHN= # If channel has to be specified
      #- FMONPB_DEV= # If device index has to be specified
    restart: unless-stopped
```

## Installation / Configuration

1. Configure Docker-Compose Environment-Variables correctly (as documented above)
2. Run Docker Container
3. Log in into Docker container
4. Run `$ pb set-key`
5. Insert your token from www.pushbullet.com/account
6. Restart Docker-Container
