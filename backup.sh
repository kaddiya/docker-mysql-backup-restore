#!/bin/bash

#!/bin/bash

if [ "${MYSQL_ENV_MYSQL_PASS}" == "**Random**" ]; then
        unset MYSQL_ENV_MYSQL_PASS
fi

#sleep 5h

MYSQL_HOST=${MYSQL_PORT_3306_TCP_ADDR:-${MYSQL_HOST}}
MYSQL_HOST=${MYSQL_PORT_1_3306_TCP_ADDR:-${MYSQL_HOST}}
MYSQL_PORT=${MYSQL_PORT_3306_TCP_PORT:-${MYSQL_PORT}}
MYSQL_PORT=${MYSQL_PORT_1_3306_TCP_PORT:-${MYSQL_PORT}}
MYSQL_USER=${MYSQL_USER:-${MYSQL_ENV_MYSQL_USER}}
MYSQL_PASS=${MYSQL_PASS:-${MYSQL_ENV_MYSQL_PASS}}

[ -z "${MYSQL_HOST}" ] && { echo "=> MYSQL_HOST cannot be empty" && exit 1; }
[ -z "${MYSQL_PORT}" ] && { echo "=> MYSQL_PORT cannot be empty" && exit 1; }
[ -z "${MYSQL_USER}" ] && { echo "=> MYSQL_USER cannot be empty" && exit 1; }
[ -z "${MYSQL_PASS}" ] && { echo "=> MYSQL_PASS cannot be empty" && exit 1; }

echo "doing a custom backup"
BACKUP_CMD="mysqldump -hdb.proof.com -P3306 -udeploy -pdeploy proof > /backup/a.sql"
MAX_BACKUPS=${MAX_BACKUPS}
BACKUP_TIME=$(date)
echo "${BACKUP_TIME}.sql"
BACKUP_NAME="${BACKUP_TIME}.sql"
echo "=> Backup started: ${BACKUP_NAME}"
$(touch /backup/a.sql)
$("mysqldump -hdb.proof.com -P3306 -udeploy -pdeploy proof > /backup/a.sql")
if $("mysqldump -hdb.proof.com -P3306 -udeploy -pdeploy proof > /backup/a.sql");then
    echo "   Backup succeeded"
else
    echo "   Backup failed"
#    rm -rf /backup/\${BACKUP_NAME}
fi

echo "=> Backup done"
