FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY ./backend .
RUN go build -v cmd/apiserver/main.go

CMD ["./main"]

