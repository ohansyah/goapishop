1. init project folder
go mod init
go mod vendor // after import package

2. environment
When it comes to creating a production-grade application, using the environment variable in the application is de facto. https://schadokar.dev/posts/go-env-ways/

Switch environment development at .env
DEVELOPMENT_TYPE = "local" || "develop" || "production"

Configure your each environment in ./configs/*.env

This project environment setting using viper. https://github.com/spf13/viper

3. middleware & Router
Middlewares are (typically) small pieces of code which take one request, do something with it, and pass it down to another middleware or the final handler. Some common use cases for middleware are request logging, header manipulation, or ResponseWriter hijacking. https://github.com/gorilla/mux.

what this middleware do :
insert log api to db
token checks

=> it comes to creating a production-grade application
Layout : single project
ENV : viper
ORM : GORM
Router : MUX

pagination
sorting
dyanmic search
degugging tools
daemon tools ??
cron
migration