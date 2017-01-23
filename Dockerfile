FROM ubuntu:trusty
MAINTAINER kaddiya <kaddiya@gmail.com>

ADD mysql-backup-restore_linux_amd64 /execute.sh
RUN chmod +x /execute.sh

RUN apt-get update && \
    apt-get install -y --no-install-recommends mysql-client && \
    apt-get install -y ca-certificates && \
    mkdir /backups

VOLUME ["/backups"]

ENTRYPOINT ["/execute.sh"]
