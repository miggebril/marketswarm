FROM golang:1.10

# Create app binary folder
RUN mkdir /app

# Copy contents
ADD marketswarm /app/marketswarm
ADD controllers /app/controllers
ADD models /app/models
ADD settings app/settings
ADD helpers app/helpers
ADD tools /app/tools
ADD helpers /app/helpers
ADD core /app/core
ADD static /app/static
ADD lib /app/lib

# Expose the application on port 8080
EXPOSE 8077

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["go", "run"]