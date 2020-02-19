FROM plugins/base:multiarch

LABEL maintainer="Weslei Juan Moser Pereira <wesleimsr@gmail.com>" \
  org.label-schema.name="Drone circleci" \
  org.label-schema.vendor="Weslei Juan Moser Pereira" \
  org.label-schema.schema-version="1.0"

ADD release/linux/amd64/drone-circleci /bin/
ENTRYPOINT ["/bin/drone-circleci"]