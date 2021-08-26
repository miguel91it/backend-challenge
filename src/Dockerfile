FROM golang:1.16

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go ./
COPY ./gateway ./gateway
COPY ./goliothMongo ./goliothMongo
COPY Makefile ./

RUN go build -o /golioth

EXPOSE 28000

CMD [ "/golioth" ]