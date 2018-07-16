FROM golang:1.10

RUN go get 

# Expose the application on port 8080
EXPOSE 8077

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["go", "run"]