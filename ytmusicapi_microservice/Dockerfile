# https://github.com/fquirin/python-ytmusicapi-server

FROM python:3.9.10-alpine3.15
MAINTAINER shur3d

# Create user
RUN adduser -D user
WORKDIR /home/user

# Install git and clone microservice
RUN apk add --no-cache --update git
#RUN git clone --branch dockerize https://github.com/SHU-red/python-ytmusicapi-server.git
RUN git clone https://github.com/fquirin/python-ytmusicapi-server.git
RUN apk del git

# Install python packages
RUN set -xe
WORKDIR /home/user/python-ytmusicapi-server
RUN python -m venv env && source env/bin/activate
RUN pip install --upgrade pip
RUN pip install fastapi
RUN pip install uvicorn
RUN pip install ytmusicapi

# Start microservice
CMD uvicorn main:api --host 0.0.0.0 --port 30010 --log-level info --reload;
