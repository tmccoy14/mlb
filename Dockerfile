ARG BASE_REGISTRY=docker.io
ARG BASE_IMAGE=golang
ARG BASE_TAG=1.18
FROM ${BASE_REGISTRY}/${BASE_IMAGE}:${BASE_TAG}

RUN apt-get update && apt-get upgrade -y

WORKDIR /app
COPY . /app

RUN go build

ENTRYPOINT ["/app/mlb"]
