docker build . -t ghcr.io/lgalkina/device-service:latest

echo access-token | docker login ghcr.io -u lgalkina --password-stdin

docker push ghcr.io/lgalkina/device-service:latest