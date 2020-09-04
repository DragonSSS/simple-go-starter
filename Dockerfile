FROM golang:1.14-alpine3.11 AS build-env

RUN apk add --no-cache git make

RUN mkdir /gostarter/

WORKDIR /gostarter/
ADD . /gostarter/

RUN make build

FROM alpine:3.8

WORKDIR /

COPY --from=build-env /gostarter/build/bin/gostarter /

EXPOSE 8080

CMD ["/gostarter"]