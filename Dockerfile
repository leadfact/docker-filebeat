FROM golang:latest

RUN mkdir /src
WORKDIR /src
COPY . /src

RUN go mod download

CMD ["go","run","main.go"]