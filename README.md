# Project Repo Title

A little introduction here...

<hr />

## **Requirements**
- Golang 1.20 or higher
- Openssl3

<hr />

## **Dependencies**
- [PymigrateDB](https://pypi.org/project/pymigratedb/)
- [Gcloud CLI](https://cloud.google.com/sdk/docs/install)

<hr />

## Setup
- To install project run ```make```.
#- It will create two new files in the project root, called: ".env" and "config.yaml"
#- Fill it with correct values to procceed with development

<hr />

## Run
- To run project execute ```make run``` into the terminal. It will start the API and serve the requests connected to the resources filled in the config files.

<hr />

## Tests
- To run and perform test cases, run the following command: ```make test```. It will begin the tests execution.

<hr />

## Accepted Methods and Content-Types

| Method | Content-Type |
|:------:|:------------:|
|POST    |application/json|
|GET     |
|OPTIONS |

## API Structure

```bash
.
├── adapters - # Adapter surface, to communicate with any client in a single interface, with standard input and output
│   ├── logging
│   │   └── logging.go
│   └── storage
│       └── storage.go
├── app - # Main application business logic directory
│   ├── errors
│   │   └── errors.go
│   ├── repository
│   │   └── .gitkeep
│   └── usecases
│       └── .gitkeep
├── clients - # Clients to implement and communicate with services and integrations provided to the API
│   ├── google
│   │   ├── logging
│   │   │   └── logging.go
│   │   └── storage
│   │       └── storage.go
│   └── iam
│       └── client.go
├── config - # Config files to make API running properly
│   ├── constants
│   │   └── constants.go
│   └── config.go
├── entity - # Entities and standard application types
│   ├── http_response.go
│   └── log.go
├── handler - # Handlers, or API entrypoints
│   └── health
│       └── health.go
├── middleware - # Middlewares to control what is being received and sent
│   ├── auth.go
│   ├── content.go
│   └── security.go
├── pkg - # Helper files
│   ├── app
│   │   └── app.go
│   ├── crypt
│   │   └── crypt.go
│   ├── helpers
│   │   ├── http.go
│   │   ├── json.go
│   │   └── utils.go
│   ├── jwt
│   │   └── jwt.go
│   └── postgres
│       └── postgres.go
├── .env.example
├── .gitignore
├── .innovation_env
├── config.example.yaml
├── go.mod
├── go.sum
├── main.go - # Golang Entrypoint and framework setup
├── Makefile
├── README.md
└── route.go - # API routing
```