go build -a -tags netgo -o release/linux/amd64/drone-circleci
docker build --rm -t wesleimp/drone-circleci -f docker/Dockerfile.linux.amd64 .