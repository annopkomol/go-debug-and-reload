FROM golang:1.15-alpine
ENV CGO_ENABLED 0

RUN apk add git
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go get github.com/cosmtrek/air

WORKDIR /src
COPY . .

RUN go mod download

ENTRYPOINT air -c .air.toml