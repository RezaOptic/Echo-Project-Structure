FROM ubuntu:16.04

RUN apt-get update && apt-get install -y ca-certificates


COPY ./app-service /

WORKDIR /

ENTRYPOINT /app-service