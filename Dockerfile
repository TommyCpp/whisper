FROM golang:1.11

COPY . /go/src/whisper/
COPY ./config/deploy_config.json /go/src/whisper/config/config.json

WORKDIR /go/src/whisper

RUN go get -u github.com/golang/dep/...
RUN dep ensure

EXPOSE 8086

RUN go build -o /app .
CMD ["/app"]