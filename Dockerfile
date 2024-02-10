FROM golang:1.21.5-alpine3.19 AS builder

WORKDIR /opt/app

COPY . .
RUN go mod tidy
RUN go build -o /opt/app/main


FROM apline:3.19.0

COPY --from=builder /opt/app/main /user/local/bin

CMD [ "/user/local/bin/main" ]
