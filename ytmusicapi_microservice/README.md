# Dockerized Microservice for YtMusicApi

## DockerHub

- https://hub.docker.com/r/shur3d/ytmusicapi_microservice

## Build

See general description at https://github.com/SHU-red/Docker

## Disclaimer

Running
- From git: Microservice from https://github.com/fquirin/python-ytmusicapi-server
- From pip: YtMusicApi from https://github.com/fquirin/python-ytmusicapi-server

special thanks to
[Florian Quirin](https://github.com/fquirin)
for his awesome work!!!

## Documentation
Please see the source repo
https://github.com/fquirin/python-ytmusicapi-server

## ToDo's
[x] Create dedicated folder for persistent settings/files

[ ] Add arm-builds?

## Docker-Compose
Example:
```
services:
  ytmusicapi_microservice:
    image: shur3d/ytmusicapi_microservice:experimental
    container_name: ytmusicapi_microservice
    environment:
      - TZ=Europe/Berlin
    volumes:
      - /mnt/tank/docker_volumes/ytmusicapi_microservice_persist:/home/user/python-ytmusicapi-server/persist
    ports:
      - 30010:30010
    restart: unless-stopped
```

## Installation
(Only necessary if you need to use authenticated search)
1. Run docker container
2. Enter python in persistent folder "persist"
3. Follow instructions, to create `headers_auth.json` like described here: https://ytmusicapi.readthedocs.io/en/latest/setup.html
4. Restart docker container
5. Test if its working by entering this in your browser: http://[YOURIP]:30010/search?q=Nevermind&maxResults=3

