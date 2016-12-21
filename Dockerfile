FROM tutum/mysql-backup
ADD backup.sh /backup.sh
RUN chmod +x /backup.sh
ADD upload.sh /upload.sh
RUN chmod +x /backup.sh
ENTRYPOINT ["/upload.sh"]
