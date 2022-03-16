FROM golang

WORKDIR /app

EXPOSE 8888

ADD go.mod go.sum ./
RUN go mod download

ADD . ./
RUN go build -o main

ENTRYPOINT [ "/app/main" ]
