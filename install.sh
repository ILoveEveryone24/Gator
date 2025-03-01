#!/bin/bash

set -e

apt-get update -y

apt-get install -y postgresql postgresql-contrib

echo "postgres:postgres" | chpasswd

service postgresql start

su - postgres -c "psql -c \"ALTER USER postgres WITH PASSWORD 'postgres';\""
su - postgres -c "psql -c \"CREATE DATABASE gator;\""
