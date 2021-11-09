FROM golang:1.17.3-bullseye

WORKDIR /golang-web-app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /server

EXPOSE 8080

CMD [ "/server" ]