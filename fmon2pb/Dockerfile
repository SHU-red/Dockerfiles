FROM python:3.9.10-alpine3.15
MAINTAINER shur3d

RUN set -xe \
    && apk add --no-cache --update \
        python3-dev \
        cargo \
        libffi-dev \
        openssl-dev \
        openrc \
    && pip install \
        pushbullet-cli \
    && apk del \
        python3-dev \
        cargo \
        libffi-dev \
        openssl-dev

RUN adduser -D user
WORKDIR /home/user

RUN mkdir -p /home/user/files
COPY ./fmon2pb /home/user
RUN chmod +x ./fmon2pb

ENTRYPOINT ./fmon2pb
