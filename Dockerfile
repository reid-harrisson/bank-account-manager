FROM golang:1.22.4

# Set destination for COPY
WORKDIR /app

# Copy the source code.
COPY . .

# Download Go modules
RUN go mod download

RUN go build -o app ./cmd/...

ENTRYPOINT [ "/app/app" ]