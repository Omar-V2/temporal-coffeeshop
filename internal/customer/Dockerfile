# Step 1: Cahce Modules
FROM golang:1.18.3-alpine3.16 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Build binary
FROM golang:1.18.3-alpine3.16 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app internal/customer/cmd/server/main.go

# Step 3: Execute binary
FROM scratch 
EXPOSE 8080
COPY --from=builder /bin/app /app
CMD ["/app"]

