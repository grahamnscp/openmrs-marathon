FROM golang:latest  AS compile
COPY data-entry-app.go /go
RUN go get "github.com/go-sql-driver/mysql"
RUN CGO_ENABLED=0 go build data-entry-app.go


FROM alpine:latest
COPY --from=compile /go/data-entry-app /

ARG MYSQL_HOST
ARG MYSQL_DATABASE
ARG MYSQL_USER
ARG MYSQL_PASSWORD

ENV MYSQL_HOST=172.17.0.2
ENV MYSQL_DATABASE=openmrs
ENV MYSQL_USER=mysqluser
ENV MYSQL_PASSWORD=mysqlpw

USER nobody:nobody
EXPOSE 8080
ENTRYPOINT ["/data-entry-app"]
