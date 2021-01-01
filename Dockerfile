# Builder
FROM golang:1.15.6-alpine3.12

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN go get github.com/rs/cors
RUN go install -v ./...

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 9090

COPY --from=builder /app/engine /app

CMD /app/engine