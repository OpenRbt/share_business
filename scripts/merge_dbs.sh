#!/bin/bash

docker cp ./merge_dbs.sql wash_postgresql:/merge_dbs.sql
docker exec -it wash_postgresql psql -U wash_bonus -f /merge_dbs.sql