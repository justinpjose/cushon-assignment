# syntax=docker/dockerfile:1

FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the repo. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
# Only copy source code
COPY ./cmd/assignment/*.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /cushon-api-binary ./cmd/assignment/*.go

# Run
CMD ["/cushon-api-binary"]