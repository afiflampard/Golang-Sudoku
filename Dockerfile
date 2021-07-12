FROM golang:latest as builder

LABEL maintainer="Afif <afiflampard@gmail.com>"
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
ENV port 8000
RUN go build
CMD [ "./TestIdn" ]