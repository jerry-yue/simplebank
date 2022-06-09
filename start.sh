#!/bin/sh

set -e

echo "check migrate file"
/app/migrate -version

echo "setup environment variable in image"
source /app/app.env

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"

exec "$@"