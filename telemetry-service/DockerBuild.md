docker build . -t ghcr.io/lygalkina/telemetry-service:latest

echo access-token | docker login ghcr.io -u lygalkina --password-stdin

docker push ghcr.io/lygalkina/telemetry-service:latest