# OAPI

Golang JSON api with sample e-commerce / oline shop project

## PROJECT LAYOUT

Mixing project and [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

### configs

Configuration default configs.
Put your .env files here.

### database

Setting database and connection

### internal

Private application and library code. 
This is the code you don't want others importing in their applications or libraries. 
Note that this layout pattern is enforced by the Go compiler itself. 
Note that you are not limited to the top level internal directory. 
You can have more than one internal directory at any level of your project tree.

### middleware

Your main middleware.

### models

NOT like MVC framework.
here we used model as representative for tabel database

### pkg

Library code that's ok to use by external applications (e.g., /pkg/mypubliclib). 
Note that the internal directory is a better way to ensure your private packages are not importable because it's enforced by Go. 
The /pkg directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. 

### queries

just like the name

### routes

set routing for each end point

### vendor

Application dependencies (managed manually or by your favorite dependency management tool like the new built-in Go Modules feature). 
The go mod vendor command will create the /vendor directory for you. 
Note that you might need to add the -mod=vendor flag to your go build command if you are not using Go 1.14 where it's on by default.
Don't commit your application dependencies if you are building a library.

## Getting started

### INIT project folder

```
go mod init
go mod vendor // after import package
```

### ENVIRONMENT

When it comes to creating a production-grade application, using the environment variable in the application is de facto. [Shubham Chadokar](https://schadokar.dev/posts/go-env-ways/).
Switch environment development at .env
DEVELOPMENT_TYPE = "local" || "develop" || "production"
Configure your each environment in ./configs/*.env
This project environment setting using [viper](https://github.com/spf13/viper).

### ROUTING

this project using [gorilla/mux](https://github.com/gorilla/mux) for routing.
implements as package routes/handle_request.json

### MIDDLEWARE

Middlewares are (typically) small pieces of code which take one request, do something with it, and pass it down to another middleware or the final handler. Some common use cases for middleware are request logging, header manipulation, or ResponseWriter hijacking.


### Log API and Queries

* API - save log every api request and responds to tb token_log (responds.ResSuccess && responds.ResErr)
* Database - print log database => main.go db.LogMode(true)


### Token

this project using [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go).
implements as package pkg/token/.there are 3 method :
* generate
* validate
* refresh

for handling token, need to reserved 2 code
* 402 - StatusPaymentRequired

token invalid -> generate new token
* 406 - StatusNotAcceptable

token expired -> refresh token

=> this project layout comes to creating a production-grade framework
[Buy Me A Coffee](https://www.buymeacoffee.com/0hans)