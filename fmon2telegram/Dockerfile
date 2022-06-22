FROM python:3.9.10-alpine3.15
MAINTAINER shur3d

RUN set -xe \
    && apk add --no-cache --update \
        python3-dev \
        cargo \
        libffi-dev \
        openssl-dev \
        openrc \
    && pip3 install \
        telegram-send \
    && apk del \
        python3-dev \
        cargo \
        libffi-dev \
        openssl-dev

RUN adduser -D user
WORKDIR /home/user

RUN mkdir -p /home/user/files
COPY ./fmon2telegram /home/user
RUN chmod +x ./fmon2telegram

ENTRYPOINT ./fmon2telegram
