#!/bin/bash

apt-get update
apt-get install -y postgresql-server-dev-17 wget unzip build-essential libssl-dev libcurl4-openssl-dev libsasl2-modules-gssapi-mit gcc libgssapi-krb5-2 libkrb5-dev

wget https://github.com/pramsey/pgsql-http/archive/refs/tags/v1.6.3.zip
unzip v1.6.3.zip
cd pgsql-http-1.6.3
make
USE_PGXS=1 make install
psql -U admin -d weather -c "CREATE EXTENSION http;"

wget https://github.com/vibhorkum/pg_background/archive/refs/tags/v1.3.zip
unzip v1.3.zip
cd pg_background-1.3
make
USE_PGXS=1 make install
psql -U admin -d weather -c "CREATE EXTENSION pg_background;"