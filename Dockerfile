FROM golang:1.16.0-alpine3.13

RUN apk update && apk add bash
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

WORKDIR /usr/app

COPY ./ ./

RUN go get -d ./...
RUN go build

EXPOSE 8080

CMD [ "./guess-pass" ]