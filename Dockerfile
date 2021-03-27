FROM golang:1.15

LABEL maintainer="Rosemary Wang"

ARG version

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-X 'main.Version=${version}'" -o hello .

EXPOSE 8001
EXPOSE 8002

# Command to run the executable
CMD ["./hello"]