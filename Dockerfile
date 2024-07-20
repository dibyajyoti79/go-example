# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
# LABEL maintainer="bookingjini<techsupport@bookingjini.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go mod tidy

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go install


# Build the Go app
RUN go build -o ./go-example

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./go-example"]