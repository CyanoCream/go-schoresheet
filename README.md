
GO-Scoresheet Readme
This repository contains the codebase for the GO-Scoresheet project. Follow the instructions below to set up the project locally.

Prerequisites
Go installed on your machine.
Swag installed (go get -u github.com/swaggo/swag/cmd/swag).

1. Ensure dependencies are in order:
    go mod tidy
2. Generate Swagger documentation for the project using:
   swag init --parseDependency --parseInternal

Accessing Swagger Documentation
After generating Swagger documentation, access it at:

http://127.0.0.1:3000/swagger/index.html#

Make sure the project is running locally before accessing Swagger documentation.