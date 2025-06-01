#! /bin/bash
set -e

# Wait for the database to be ready
until psql -h db -U postgres -d postgres -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

psql -h db -U postgres -d postgres -f tables.sql

exec go run main.go