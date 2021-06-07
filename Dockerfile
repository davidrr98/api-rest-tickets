FROM ubuntu:19.10

RUN mkdir /code

WORKDIR /code

RUN apt-get update

RUN apt-get install -y nginx

