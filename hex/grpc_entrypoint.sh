#!/bin/sh

set -e
COMMAND=$@

# Try to connect to the database every few secons for ten tries
maxTries=10
echo 'Waiting for database to become available...'
while [ "$maxTries" -gt 0 ] && ! mysql -h "$MYSQL_HOST" -P "$MYSQL_PORT" -u "$MYSQL_USER" -p "$MYSQL_PASSWORD" -e 'SHOW TABLES'; do
    maxTries=$(($maxTries - 1))
    sleep 3
done
echo
if [ "$maxTries" -le 0 ]; then
    echo >&2 'error: unable to contact mysql after 10 tries'
    exit 1
fi

exec $COMMAND