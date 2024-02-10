FROM golang:1.21.5-alpine3.19 AS builder


RUN apk update && apk add git && rm -rf /var/cache/apk/*

WORKDIR /opt/app

COPY ./app .
COPY ./config .
COPY ./cmd .
COPY ./internal .
COPY ./utils .
COPY ./.air.toml .
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download
RUN go build -o /opt/app/main


FROM apline:3.19.0

COPY --from=builder /opt/app/main /user/local/bin

CMD [ "/user/local/bin/main" ]
