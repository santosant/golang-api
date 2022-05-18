# Go-API

Implementing an API with Golang and Postgres

In this application will defined in 4 packages, which are:

### DB

Responsible for Database related settings

### Configs

Responsible for taking these settings to the application.

### Handlers

Responsible for receiving and handling API calls.

### Models

Responsible for integration with DB, receives requests from Database and creates tasks in Postgres

## How to Run

You just need run the command go run main.go in the root of project

to test the endpoints e connect to a real postgres db, was used Docker Container with postegress

```docker
  docker run -d --name api_todo -p 5432:5432 -e POSTEGRES_PASSWORD=1234 postgres:13.5
```

just run this command enter on postegress and create database and table.
