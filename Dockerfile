FROM golang:1.18-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server main.go

FROM alpine:3.10
WORKDIR /app
COPY --from=build /app/server .
COPY --from=build /app/env.yaml .


EXPOSE 8000

CMD ["/app/server"]

