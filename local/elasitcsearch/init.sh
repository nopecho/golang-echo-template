#!/bin/bash

LOCAL_ES_URL=localhost:9200

function init() {
    while true; do
        if curl --silent --fail "$1"/_cluster/health; then
            echo "Elasticsearch is up and running"
            break
        else
            echo "Elasticsearch is not running"
        fi
        sleep 3
    done
}

init $LOCAL_ES_URL &
/bin/tini -- /usr/local/bin/docker-entrypoint.sh