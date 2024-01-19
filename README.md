# udacity-crm-backend-go

This is the final project of the Udacity Go Language (Golang) course, which will be submitted.
It is a backend completely written in Go, offering a few endpoints in order to process, save and change customer info.
A small API doc is also included at the root path of the server.

# Installation

The installation of the project is quite straight forward.
Pull the repository and make sure that go and the gorilla/mux package is installed:

GoLang: 
https://go.dev/doc/install
```
go get -u github.com/gorilla/mux
```
# Launching the application
When everything is set up, simply go into the root directory of the project and type:
```
go run main.go
```
The server should start with a message on the specified port _(default :3000)_

# Usage
The application is a simple API-Backend Server with no direct frontend for the functionality, meaning using a tool like Postman or cUrl is advised.

The application has 3 users prefilled in the db for testing purposes.

As a first example, try to get all users currently in the DB:

```
curl localhost:3000/customers
```

This command sends a simple GET request for all users and retrieves them in JSON format.

To see all possible API calls, refer to the small [API-Doc](http://localhost:3000/) at the root path. _(Server needs to be running)_

__Have fun Testing!__