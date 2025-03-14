#!/bin/bash

sed -i '/^#logging_collector/c\logging_collector = off' var/lib/postgresql/data/postgresql.conf
sed -i '/^#log_statement/c\log_statement = 'none'' var/lib/postgresql/data/postgresql.conf
sed -i '/^#log_min_duration_statement/c\log_min_duration_statement = -1' var/lib/postgresql/data/postgresql.conf
