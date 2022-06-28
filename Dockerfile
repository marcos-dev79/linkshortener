FROM golang:1.13.1

WORKDIR /go/src/app
COPY . .

RUN go install linkshortener
CMD ["linkshortener"]