FROM golang:1.25-rc

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

CMD [ "air" ]