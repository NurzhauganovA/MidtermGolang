# Final Project by Golang

## Set UP Project

docker pull PostgreSQL -> `docker pull postgres`

docker run PostgreSQL -> `docker run --name=online-store-db -e POSTGRES_PASSWORD='********' -p 5436:5432 -d --rm postgres` (password is hashing!)

docker check containers -> `docker ps`

create migrate for schema of postgres -> `migrate create -ext sql -dir ./schema -seq init`

up migrated schema -> `migrate -path ./schema -database postgres://postgres:'********'@localhost:5436/postgres?sslmode=disable up`

check upped container of postgres -> `docker exec -it ******* /bin/bash` (******* -> docker container name)

sign in postgres container -> `psql -U postgres`

check all schemes -> `\d`

![image](https://user-images.githubusercontent.com/90419990/218040724-b1e82320-e7ef-4f34-a598-a046e35679fa.png)

down migrated schema -> `migrate -path ./schema -database postgres://postgres:'********'@localhost:5436/postgres?sslmode=disable down`

update schema version -> `update schema_migrations set version='000001', dirty=false;`



## Progress Report #1

We started the project slowly. We thought we would deliver on the set day, but as you extended the deadline for the report, we decided to finish one part of the project entirely. 

***So what did we do:***

***We have completely completed the implementation of the main functions of simple implementations. Until the next report, we will try to finish the midterm project already. We implement the authorization function, we are going to implement authentication through JWT tokens. Then we have 1 function left, this is CRUD. In fact, for 1 midterm project, it's enough just to show the products from the database. But we will try and create these products through our website on the Admin page.***

`cmd` -> here is the main file itself that needs to be run.

`configs` -> here is the database configuration, more precisely connect with PostgreSQL. But we deleted this file from GitHub, because there is secret information (passwords) there.

`endpoint` -> here is the overall structure of our tables.

`package/handler` -> here are mainly routers, and the controller implementation system itself.

`package/repository` -> interfaces have been created here, so far the registration system has been implemented. The function of creating a new user has been created.

`package/service` -> interfaces are also created here. A password hashing function has been created. You can see the implementation of the `generatePasswordHash` functions

`schema` -> PostgreSQL files are created here, namely a file for creating tables and a file for deleting tables.

There is also an `ERD` you can look at the table, I'm sure we will expand it.
