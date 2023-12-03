# FROM golang:1.20.2 as builder

# WORKDIR /todolist/

# COPY . .

# RUN CGO_ENABLED=0 go build -o todolist-microservice /todolist/main/main.go

# FROM alpine:latest

# WORKDIR /todolist

# COPY --from=builder /todolist/ /todolist//

# EXPOSE 8080

# CMD ./todolist



# cat Dockerfile
FROM mysql:latest


RUN chown -R mysql:root /var/lib/mysql/

ARG MYSQL_DATABASE
# ARG MYSQL_USER
# ARG MYSQL_PASSWORD
ARG MYSQL_ROOT_PASSWORD

ENV MYSQL_DATABASE=dacn2
# ENV MYSQL_USER=buiduchieu
# ENV MYSQL_PASSWORD=root
ENV MYSQL_ROOT_PASSWORD=root

ADD data.sql /etc/mysql/data.sql

RUN sed -i 's/MYSQL_DATABASE/'$MYSQL_DATABASE'/g' /etc/mysql/data.sql
RUN cp /etc/mysql/data.sql /docker-entrypoint-initdb.d

EXPOSE 3306