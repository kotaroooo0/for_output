#!/bin/ash

set -e

until nc -z redis 6379; do
  >&2 echo "redis is unavailable - sleeping"
  sleep 3
done
until nc -z mysql 3306; do
  >&2 echo "mysql is unavailable - sleeping"
  sleep 3
done
>&2 echo "mysql is up - executing command"

exec $@
