FROM golang:latest

WORKDIR /usr/src/app

COPY ./ecoridegolang /usr/src/app/
COPY ./.env /usr/src/app/

RUN go mod tidy

CMD ["go", "run", "main.go"]