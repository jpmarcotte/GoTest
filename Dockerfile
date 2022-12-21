FROM golang:1.19-alpine

RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
COPY src/ ./src/
COPY data/ ./data/

RUN file="$(ls -l ./)" && echo "$file"
RUN file="$(ls -l ./src/)" && echo "$file"
RUN go build -o ./webserver

EXPOSE ${PORT}

CMD [ "./webserver" ]