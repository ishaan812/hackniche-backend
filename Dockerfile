FROM golang:1.19-alpine

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download 

COPY . .
RUN go build -o /weber

EXPOSE 9000

CMD ["/weber"]

