# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9.1

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/maknahar/go-web-skeleton/

# Build the application inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "glide".)
RUN go install github.com/maknahar/go-web-skeleton/

# Run the application by default when the container starts.
ENTRYPOINT /go/bin/go-web-skelton

# Document that the service listens on port 8000.
EXPOSE 8000
