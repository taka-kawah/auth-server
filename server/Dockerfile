FROM golang:1.24.3-bookworm

RUN apt update
RUN apt install -y postgresql-client


WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

CMD ["./scripts/init.sh"]