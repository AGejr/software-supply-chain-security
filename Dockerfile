# Start from a lightweight image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the current directory
COPY ./main /app/main

# Document that the service listens on port 8080
EXPOSE 8080

# Run the binary
CMD ["./main"]

