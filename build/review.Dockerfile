FROM golang:alpine as build

COPY . /project

WORKDIR /project

RUN apk add make && make build TARGET=review

#========================================

FROM alpine:latest

COPY --from=build /project/bin/review /bin/

RUN apk update && apk add bash

WORKDIR /project

RUN apk add --no-cache tzdata
ENV TZ="Europe/Moscow"

CMD ["review"]
