FROM golang

WORKDIR /app

EXPOSE 8888

ADD go.mod go.sum ./
RUN go mod download

ADD main.go ./
ADD event.go ./
ADD nature.go ./
ADD operation.go ./
ADD query.go ./
RUN go build -o main

ENTRYPOINT [ "/app/main" ]
