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

# Expose port 8077
EXPOSE 8077

ENTRYPOINT /app/marketswarm