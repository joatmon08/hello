FROM golang:1.12

LABEL maintainer="Rosemary Wang"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o hello .

EXPOSE 8001
EXPOSE 8002

# Command to run the executable
CMD ["./hello"]