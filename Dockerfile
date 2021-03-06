FROM golang:1.16.0-alpine3.13 as build

WORKDIR /app

COPY ./Search-Service .

RUN go mod download
RUN go build -o /searchservice

FROM alpine:3.13
COPY --from=build /app/config/ ./config/
COPY --from=build /searchservice /searchservice

ENTRYPOINT  ["/searchservice"]