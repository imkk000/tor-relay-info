# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/tor-relay-info .

FROM gcr.io/distroless/static:nonroot
COPY --from=build /out/tor-relay-info /tor-relay-info
ENTRYPOINT ["/tor-relay-info"]
