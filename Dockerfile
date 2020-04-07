FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN apk add \
    build-base
RUN go build -o main .

FROM alpine
COPY --from=builder /build/main /app/
COPY --from=builder /build/wait-for-postgres.sh /app/
RUN apk add \    
    postgresql-client
WORKDIR /app
CMD ["./wait-for-postgres.sh", "pgsql", "./main"]
