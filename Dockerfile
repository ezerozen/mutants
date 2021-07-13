FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN go build -o /api

EXPOSE 8080

CMD [ "/api", "run"]
