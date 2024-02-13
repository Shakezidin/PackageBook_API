FROM golang:1-alpine

LABEL maintainer="Shaikhzidhin <sinuzidin@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

ENV APIPORT="7070"
ENV GRPCUSERPORT="9090"
ENV GRPCADMINPORT="8080"
ENV GRPCCORDINATORPORT="7000"
EXPOSE $APIPORT $GRPCUSERPORT $GRPCADMINPORT $GRPCCORDINATORPORT

CMD ["./main"]

