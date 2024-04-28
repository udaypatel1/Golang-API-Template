# CRUD API in Go Template

This is a simple template that can store information in an in-memory database, written in Go

It can be extensible and customized further to one's liking

Please feel free to fork or clone

## Setup

`go run main.go`

Server will be live by default on `localhost:8080`

### List All Items

`curl localhost:8080/items`

### Add an Item

`curl --request POST --data '{"id": 1, "name": "Item 1"}' localhost:8080/items/add`

