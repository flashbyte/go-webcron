FROM golang:1.19.2 as build

WORKDIR /app

RUN apt-get update && apt-get install -y upx

COPY cron.go .

ENV CGO_ENABLED=0
RUN go build -ldflags "-s -w" cron.go

RUN upx cron


FROM scratch

WORKDIR /

COPY --from=build /app/cron /cron

ENTRYPOINT ["/cron"]
