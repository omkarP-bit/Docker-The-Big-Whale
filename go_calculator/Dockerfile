FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod .
COPY calculator.go .

RUN CGO_ENABLED=0 go build -o calculator .

FROM scratch

COPY --from=build /app/calculator /calculator

ENTRYPOINT ["/calculator"]