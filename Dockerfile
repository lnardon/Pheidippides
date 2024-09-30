FROM golang:1.23.1-alpine3.19 AS build-backend
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:3.19 AS runner
WORKDIR /usr/src/app
COPY --from=build-backend /usr/src/app/main /usr/src/app/main

EXPOSE 8080
CMD ["./main"]