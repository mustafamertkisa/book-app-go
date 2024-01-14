#!/bin/bash

docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest
sleep 3
echo "Postgresql starting..."

docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE bookapp"
sleep 3
echo "Database bookapp created"

docker exec -it postgres-test psql -U postgres -d bookapp -c "
create table if not exists books
(
    id bigserial not null primary key,
    name varchar(255) not null,
    pages integer not null,
    author varchar(255) not null
)
"
sleep 3
echo "Table books created"