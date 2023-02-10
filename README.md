# Final Project by Golang

## Progress

docker pull PostgreSQL -> `docker pull postgres`

docker run PostgreSQL -> `docker run --name=online-store-db -e POSTGRES_PASSWORD='********' -p 5436:5432 d --rm postgres` (password is hashing!)

docker check containers -> `docker ps`

create migrate for schema of postgres -> `migrate create -ext sql -dir ./schema -seq init`

up migrated schema -> `migrate -path ./schema -database postgres://postgres:'********'@localhost:5436/postgres?sslmode=disable up`

check upped container of postgres -> `docker exec -it ******* /bin/bash` (******* -> docker container name)

sign in postgres container -> `psql -U postgres`

check all schemes -> `\d`

down migrated schema -> `migrate -path ./schema -database postgres://postgres:'********'@localhost:5436/postgres?sslmode=disable down`

update schema version -> `update schema_migrations set version='000001', dirty=false;`
