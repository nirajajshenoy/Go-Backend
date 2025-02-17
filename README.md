
## Go Fiber Backend API

#  Overview

This is a simple backend API built using the Fiber framework for Go. It includes basic route handling for GET, POST, and DELETE requests.

# Features

GET /hello: Calls GetController from the controllers package.

POST /hello: Calls PostController from the controllers package.

DELETE /hello: Calls DeleteController from the controllers package.

GET /: Returns a message: Get All Data.

POST /: Returns a message: To post a data.

GET /api: Returns a message: Get particular Data.

DELETE /api: Returns a message: Delete a data.

# Installation

## Clone the repository:

git clone https://github.com/nirajajshenoy/Go-Backend.git

## Change into the project directory:

cd Go-Backend

## Install dependencies:

go mod tidy

## Running the Server

Run the following command to start the server:

go run main.go

The server will start on http://localhost:8080.

## API Endpoints

/hello

GET /hello → Calls GetController

POST /hello → Calls PostController

DELETE /hello → Calls DeleteController

/

GET / → Returns Get All Data

POST / → Returns To post a data

/api

GET /api → Returns Get particular Data

DELETE /api → Returns Delete a data
