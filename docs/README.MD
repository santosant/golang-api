# Go-API

Implementando uma API com Golang e PostgreSQL

Esta aplicação serão definidas em 4 packages, que são eles:

### DB

Responsável pelas configurações relacionado ao Database

### Configs

Responsável por levar estas configurações para a aplicação.

### Handlers

Responsável por receber e tratar as chamadas da API.

### Models

Responsável com a integração com o DB, recebe as requestes do Banco e cria as tarefas no Postgres

## How to Run

You just need run the command go run main.go in the root of project

to test the endpoints e connect to a real postgres db, was used Docker Container with postegress

```docker
  docker run -d --name api_todo -p 5432:5432 -e POSTEGRES_PASSWORD=1234 postgres:13.5
```

just run this command enter on postegress and create database and table.
