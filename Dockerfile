FROM ubuntu:latest

COPY bin/webserver .
ENTRYPOINT ./webserver