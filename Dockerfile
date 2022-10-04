# One of the most common ways to accidentally build large images is to do the actual 
# program compilation as part of the construction of the application container image.
# With multistage builds, rather than producing a single image, a Docker file can actually
# produce multiple images. Each image is considered a stage. Artifacts can be copied from
# preceding stages to the current stage.

# Building a container image using multistage builds can reduce your final container image size
# by hundreds of megabytes and thus dramatically speed up your deployment times, since generally,
# deployment latency is gated on network performance.

# STAGE1: BUILD. produces build image which contains the Go compiler and source code for the program
FROM golang:1.17-alpine AS build

# Specify the directory inside the image in which all commands will run
WORKDIR /usr/src/app

# Copy package files and install dependencies
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

# Copy all of the app files into the image
COPY . .

# Specify set of variables that the build script expects

# Do the build.
RUN CGO_ENABLED=0
RUN go install ./cmd

# STAGE 2: DEPLOYMENT. produces deployment image, which simply contains the compiled binary.
FROM alpine
COPY --from=build /go/bin/cmd /
# The default command to run when starting the container
# CMD ["/cmd"]
ENTRYPOINT [ "/browse" ]


