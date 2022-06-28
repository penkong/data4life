# # Use the offical Golang image to create a build artifact.
# # This is based on Debian and sets the GOPATH to /go.
# # https://hub.docker.com/_/golang
# FROM golang:1.18 as builder

# # Copy local code to the container image.
# WORKDIR /go/app
# COPY . .

# # Build the command inside the container.
# # (You may fetch or manage dependencies here,
# # either manually or with a tool like "godep".)
# RUN CGO_ENABLED=0 GOOS=linux go build -v -o app cmd/main.go

# # Use a Docker multi-stage build to create a lean production image.
# # https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
# FROM alpine
# RUN apk add --no-cache ca-certificates

# # Copy the binary to the production image from the builder stage.
# COPY --from=builder /go/app/app /app

# # Run the web service on container startup.
# CMD ["/app"]

FROM golang as build-go
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app .

FROM alpine:latest
RUN addgroup -S app && adduser -S app -G app
USER cloud-run-example
WORKDIR /home/app
COPY --from=build-go /bin/app ./
EXPOSE 3000
ENTRYPOINT ["./app"]
