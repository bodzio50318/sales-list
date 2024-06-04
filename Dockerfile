FROM golang:alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -x -o /sales

EXPOSE 8080

CMD [ "/sales" ]