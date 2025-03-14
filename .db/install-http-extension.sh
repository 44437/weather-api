#!/bin/bash

apt-get update
apt-get install -y postgresql-server-dev-17 wget unzip build-essential libssl-dev libcurl4-openssl-dev
wget https://github.com/pramsey/pgsql-http/archive/refs/tags/v1.6.3.zip
unzip v1.6.3.zip
cd pgsql-http-1.6.3
make
USE_PGXS=1 make install
psql -U admin -d weather -c "CREATE EXTENSION http;"