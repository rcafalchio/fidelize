# widget-go-spa-api
API using Golang, MongoDB and JWT Authentication.

# Dependencies
- go^1.8
- MongoDB^3.4.2

# How to run
- Download or clone this repository to your GO $WorkSpace
- Change directory into the project and Install project dependencies
```sh
$ go get ./...
```
- Open settings.config on common folder and change variables, depending on your configuration
- Build the application
```sh
$ go build
```
- Run the application (Make sure the mongodb is running)

### Windows
```sh
$ go run main.go
```
### MAC and Linux
```sh
$ ./widgets-spa-go-api
```

- The app will start at port 9000


# Endpoints

- GET /authenticate http://localhost:9000/authenticate
- GET /users http://localhost:9000/users
- GET /users/:id http://localhost:9000/users/:id
- GET /widgets http://localhost:9000/widgets
- GET /widgets/:id http://localhost:9000/widgets/:id
- POST /widgets for creating new widgets http://localhost:9000/widgets
- PUT /widgets/:id for updating existing widgets http://localhost:9000/widgets/:id

# How to use

## AUTHENTICATE
- Make a basic get request to get a token

```sh
$ curl http://localhost:9000/authenticate
```

## GET USERS
- Add 'Authorization: Bearer {{TOKEN}} to Header request
- {{TOKEN}} is the value returned by localhost:9000/authenticate
- Do a GET Request
```sh
$ http://localhost:9000/users
```

## GET USER BY ID
- Add 'Authorization: Bearer {{TOKEN}} to Header request
- {{TOKEN}} is the value returned by localhost:9000/authenticate
- Do a Get request
- Change :id for the User id
```sh
$  http://localhost:9000/users/:id
```

## GET WIDGETS
- Add 'Authorization: Bearer {{TOKEN}} to Header request
- {{TOKEN}} is the value returned by localhost:9000/authenticate
- Do a GET request
```sh
$  http://localhost:9000/widgets 
```

## GET WIDGET BY ID
- Add 'Authorization: Bearer {{TOKEN}} to Header request
- {{TOKEN}} is the value returned by localhost:9000/authenticate
- Do a GET request
- Change :id for the Widget id
```sh
$  http://localhost:9000/widgets/:id 
```

## CREATE WIDGET
- Add 'Authorization: Bearer {{TOKEN}} to Header request
- {{TOKEN}} is the value returned by localhost:9000/authenticate
- Do a post request Adding a json to the body
-Example JSON

```sh
{"name":"WIDGET TEST POST ","color":"RED","price":"10","inventory":5,"melts":true}
```
```sh
$  http://localhost:9000/widgets
```

## UPDATE WIDGET
- Add 'Authorization: Bearer {{TOKEN}} to Header request
- {{TOKEN}} is the value returned by localhost:9000/authenticate
 -Do a put request Adding a json to the body
- Change :id for the Widget id

```sh
{"name":"UPDATING TEST","color":"yello","price":"40","inventory":2,"melts":false}
```
```sh
$  http://localhost:9000/widgets/:id
```
