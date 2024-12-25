#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path db/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"