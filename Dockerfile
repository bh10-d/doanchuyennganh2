FROM golang:1.20.2 as builder

WORKDIR /todolist/

COPY . .

RUN CGO_ENABLED=0 go build -o todolist-microservice /todolist/main/main.go

FROM alpine:latest

WORKDIR /todolist

COPY --from=builder /todolist/ /todolist//

EXPOSE 8080

CMD ./todolist