FROM golang:alpine as build

COPY . /project

WORKDIR /project

RUN apk add make && make build

#========================================

FROM alpine:latest

COPY --from=build /project/bin/app /bin/

RUN apk update && apk add bash

WORKDIR /project

RUN apk add --no-cache tzdata
ENV TZ="Europe/Moscow"

CMD ["app"]
