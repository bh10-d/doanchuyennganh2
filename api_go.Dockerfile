FROM golang:latest as builder

WORKDIR /go-microservice

COPY ./golang .

RUN CGO_ENABLED=0 go build -o microservice /go-microservice/main/main.go

FROM alpine:latest

WORKDIR /go-microservice

COPY --from=builder /go-microservice/ /go-microservice/

EXPOSE 8080

CMD ./microservice
# CMD ./todolist


# # Specify the base image for the go app.
# FROM golang:1.20.2
# # Specify that we now need to execute any commands in this directory.
# WORKDIR /todolist
# # Copy everything from this project into the filesystem of the container.
# COPY . .
# # Obtain the package needed to run redis commands. Alternatively use GO Modules.
# # RUN go get -u github.com/go-sql-driver/mysql
# RUN go get -u github.com/gin-gonic/gin
# RUN go get -u gorm.io/gorm
# RUN go get -u gorm.io/driver/mysql
# # Compile the binary exe for our app.
# RUN go build -o main .
# # Start the application.
# CMD ["./main/"]