FROM golang:1.20

ADD ./greet/greet greet

ENTRYPOINT exec go run greet

EXPOSE 8888