FROM ubuntu:trusty
MAINTAINER kaddiya <kaddiya@gmail.com>


RUN apt-get update && \
    apt-get install -y --no-install-recommends mysql-client && \
    mkdir /backup

VOLUME ["/backup"]
