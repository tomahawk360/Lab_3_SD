FROM golang:1.21.3 AS builder

ARG var

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

CMD if [ "$var" = "broker" ]; then \
    cd /app/broker; \
    go run broker.go; \
    if [ "$var" = "vanguardia" ]; then \
    cd /app/vanguardia; \
    go run vanguard.go; \
    elif [ "$var" = "f1" ]; then \
    cd /app/fulcrum1; \
    go run fulcrum1.go; \
    elif [ "$var" = "f2" ]; then \
    cd /app/fulcrum2; \
    go run fulcrum2.go; \
    elif [ "$var" = "f3" ]; then \
    cd /app/fulcrum3; \
    go run fulcrum3.go; \
    elif [ "$var" = "i1" ]; then \
    cd /app/caiatl; \
    go run informante.go; \
    elif [ "$var" = "i2" ]; then \
    cd /app/osiris; \
    go run informante.go; \
    else \
    echo "Invalid value for 'var'"; \
    fi