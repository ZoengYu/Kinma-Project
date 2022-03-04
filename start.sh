#!/bin/sh

set -e

echo "run db migration"
#migrate require the $DB_SOURCE variable to connect to db
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"