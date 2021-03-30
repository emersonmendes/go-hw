#!/bin/bash

set -e

psql --quiet -c "ALTER USER postgres PASSWORD 'postgres'"
psql -d hw_db -U postgres -f /usr/src/db.sql