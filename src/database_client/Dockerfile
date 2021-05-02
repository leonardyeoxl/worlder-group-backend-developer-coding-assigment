FROM common-grpc-alpine:latest

COPY go.mod  /app/go.mod
RUN go mod download

# COPY . .

# ENTRYPOINT [ "go", "run", "db-client.go" ]