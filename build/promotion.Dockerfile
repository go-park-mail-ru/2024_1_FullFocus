FROM golang:alpine as build

COPY . /project

WORKDIR /project

RUN apk add make && make build TARGET=promotion

#========================================

FROM alpine:latest

COPY --from=build /project/bin/promotion /bin/

RUN apk update && apk add bash

WORKDIR /project

RUN apk add --no-cache tzdata
ENV TZ="Europe/Moscow"

CMD ["promotion"]
