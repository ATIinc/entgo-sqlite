#!/bin/bash

# https://github.com/coleifer/sqlite-web

# Run this on the host, not in the devcontainer

PROJ_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

port="8012"

db="$PROJ_DIR/example.db"
db_dir=$(dirname $db)
db_file=$(basename $db)
echo "Starting on: http://127.0.0.1:$port connecting to $db"
docker run -i -d --rm -p "$port:$port" -v $db_dir:/data -e SQLITE_DATABASE="$db_file" coleifer/sqlite-web sqlite_web -H 0.0.0.0 -p $port -r -x $db_file