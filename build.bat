go build -a -tags netgo -o release/windows/amd64/drone-circleci.exe
docker build --rm -t wesleimp/drone-circleci -f docker/Dockerfile.windows.1903.amd64 .