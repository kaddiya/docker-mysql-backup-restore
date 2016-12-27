FROM ubuntu:trusty
MAINTAINER kaddiya <kaddiya@gmail.com>

RUN apt-get update && \
    apt-get install -y --no-install-recommends mysql-client && \
    apt-get install -y ca-certificates && \
    mkdir /backups

VOLUME ["/backups"]

ENTRYPOINT ["/execute.sh"]
