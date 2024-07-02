#syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build  -o /goapp ./cmd

EXPOSE 8080

CMD ["/goapp"]


