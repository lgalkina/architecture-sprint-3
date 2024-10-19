docker build . -t ghcr.io/lgalkina/telemetry-service:latest

echo access-token | docker login ghcr.io -u lgalkina --password-stdin

docker push ghcr.io/lgalkina/telemetry-service:latest