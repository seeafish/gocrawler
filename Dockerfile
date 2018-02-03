FROM golang:1.9

WORKDIR /go/src/app
COPY . .

WORKDIR crawler

RUN go test -cover
RUN go install

CMD ["crawler"]