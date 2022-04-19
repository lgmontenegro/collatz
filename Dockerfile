FROM golang:1.14-alpine3.11 as stage

WORKDIR /app
COPY ./src .
RUN env CGO_ENABLED=0 go build -tags netgo -ldflags '-w -s -extldflags "-static"' -o collatz

FROM scratch
COPY --from=stage /app/collatz .
CMD ["./collatz"]