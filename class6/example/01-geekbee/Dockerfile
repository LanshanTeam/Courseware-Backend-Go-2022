FROM golang:1.18-alpine AS build

WORKDIR /go/src/github.com/gocybee/Courseware-2022

COPY main.go .

RUN go build -o server main.go

FROM build AS dev

RUN apk update \
    && apk add git

COPY website website

CMD ["go", "run", "main.go"]

FROM alpine:3.12

WORKDIR webapp

RUN apk --no-cache add ca-certificates

COPY website website
COPY --from=build /go/src/github.com/gocybee/Courseware-2022/server server

EXPOSE 8077

CMD ["/webapp/server"]