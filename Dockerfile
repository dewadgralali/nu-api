FROM golang AS build

WORKDIR /go/src/nu

ADD . .

RUN go get -v ./...
RUN go build

FROM alpine

WORKDIR /usr/bin

COPY --from=build /go/src/nu/nu nu
RUN chmod a+x nu

CMD ["nu", "serve"]
