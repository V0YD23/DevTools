# Use a Go base image
FROM golang:1.22.2-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules and Go sum files to the container (if you have them)
COPY go.mod ./
# COPY go.sum ./

# If the go.sum file is missing, you can skip the copy but still run `go mod tidy` to generate it
RUN if [ -f go.sum ]; then go mod tidy; else echo "go.sum not found, skipping go mod tidy"; fi

# Copy the rest of the application code (including ops folder) to the container
COPY . .

# Install make utility
RUN apk add --no-cache make

# Run the default target (build, run, or test) depending on the command in Docker
CMD ["make", "run"]
