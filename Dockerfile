FROM golang:alpine

WORKDIR /go/src/gitlab.com/idoko/letterpress

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /usr/bin/letterpress ./cmd/api

EXPOSE 8080

ENTRYPOINT ["/usr/bin/letterpress"]
