FROM golang:1.21

WORKDIR /backend

COPY go.mod go.sum ./

RUN ["go", "mod", "download"]

COPY . .

RUN ["go", "build", "-o", "./build/backend", "./cmd/houz"]

CMD ["./build/backend"]
