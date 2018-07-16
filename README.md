## Create shareable image
docker build -t market-image .

## Make executable update script
chmod +x update.sh

## Run image
docker run -it --rm --name market-instance -p 8077:8077 -v /app/marketswarm:/go/src/marketswarm -w /go/src/marketswarm market-image
