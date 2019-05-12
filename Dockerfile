FROM golang AS build

WORKDIR /go/src/github.com/dewadg/nu-api

ADD . .

RUN go get -v ./...
RUN CGO_ENABLED=0 go build

FROM alpine

WORKDIR /usr/local/bin

COPY --from=build /go/src/github.com/dewadg/nu-api/nu-api ./nu
RUN chmod +x ./nu

CMD ["nu", "serve"]
