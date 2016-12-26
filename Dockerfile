FROM ubuntu:trusty
MAINTAINER kaddiya <kaddiya@gmail.com>

ADD docker-mysql-backup-restore_linux_amd64 /start.sh
RUN chmod +x /start.sh

RUN apt-get update && \
    apt-get install -y --no-install-recommends mysql-client && \
    mkdir /backups

RUN apt-get install -y ca-certificates

VOLUME ["/backups"]

ENTRYPOINT ["/start.sh"]
