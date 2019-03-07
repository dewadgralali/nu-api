FROM golang AS build

WORKDIR /go/src/nu

ADD . .

RUN go get -v ./...
RUN CGO_ENABLED=0 go build

FROM alpine

WORKDIR /usr/local/bin

COPY --from=build /go/src/nu/nu .
RUN chmod +x ./nu

CMD ["nu", "serve"]
