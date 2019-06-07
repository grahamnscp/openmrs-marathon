FROM alpine:latest

RUN apk add mysql-client && apk add curl && apk add bash
COPY * /
CMD ["/daemon-wrapper.sh"]
