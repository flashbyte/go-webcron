FROM golang:1.19.2 as build

WORKDIR /app

COPY cron.go .

RUN go build -v cron.go

FROM scratch 

WORKDIR /

COPY --from=build /app/cron /cron

ENTRYPOINT ["/cron"]
