# Go Version
FROM golang:latest

# Environment variables which CompileDaemon requires to run
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# Basic setup of the container
RUN mkdir /app

WORKDIR /app

COPY ./src ./
RUN go mod download

# Get CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

RUN apt-get update && apt-get install -y unzip

EXPOSE 3000
EXPOSE 50051
# The build flag sets how to build after a change has been detected in the source code
# The command flag sets how to run the app after it has been built
ENTRYPOINT CompileDaemon -build="go build -o api" -command="./api"