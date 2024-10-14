FROM golang:latest
WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -v -o server
CMD ["/app/server"]