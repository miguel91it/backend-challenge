FROM golang:1.16

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY *.go ./
COPY Makefile ./

RUN ls

RUN go build -o /devices

CMD [ "/devices"]