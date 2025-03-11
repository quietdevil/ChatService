#!/bin/bash

source .env

export LOCAL_MIGRATION_DSN="host=pg port=$PG_PORT dbname=$PG_DATABASE_NAME user=$PG_USER password=$PG_PASSWORD"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${LOCAL_MIGRATION_DSN}" up -v




