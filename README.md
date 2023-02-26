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




## Progress Report #2

So that's what we added in the second report. Last time we added completely only the registration system with `Name`, `Surname`, `Username`, `Email` and `Password`. Now we have added authorization via JWT tokens to this system. This is implemented in the `Sign-In` method. We also created a `generatePasswordHash` function through which we can get a hashed password from the database.

Also, if you noticed, last time we run all this every time on the terminal with `Docker`. Starting first `PostgreSQL`, then `Golang` server. And now, we have implemented all this through `docker-compose`, creating `DockerFile` and `docker-compose.yml`.

We also decided to add `React App`. If we assumed correctly, you have removed the frontend part from the project. But we decided to add it anyway, since we can actually see what is happening under the hood of the program system.
So far, 3 pages have been added. The main 2 is the Registration page and the Authorization page. So far, we have had an error with `CorsHeaders`, but when we launched `React App`, linking with `Django`, we solved this problem very easily. This time on `Golang` we hope to solve it quickly too. That's all.


## Progress Report #3

So, what we did during the week, especially nothing) We didn't have enough time for the project, due to the fact that we are preparing for an internship in the same company as a whole team. Okay, not the point, after all, we added something interesting.

Added `middleware`. In order to check the corresponding user token and respond to the request judging by this token.

What does it mean? We have created such a system, in the future the program can check the token and if this token is correct, it will output us, for example, a list of all products. It's kind of like permissions. If the client who requests the url `http://localhost:8000/api/categories` there is a token (this means he logged in), gets the correct answer.
