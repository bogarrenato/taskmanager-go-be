FROM golang:latest

# Download watcher on startup
RUN go get github.com/canthefason/go-watcher/cmd/watcher

# Specify the working directory with the name of the project
WORKDIR /Users/renatobogar/Documents/GitHub/taskmanager-go-be

# Copy all the files to the container
COPY . .

# Run the watcher
ENTRYPOINT /go/bin/watcher