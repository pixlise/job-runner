FROM golang:1.24-alpine AS builder

RUN apk add --no-cache build-base

COPY ./runner /build
ENV CGO_ENABLED=1

RUN cd /build && GOOS=linux GOARCH=amd64 go build -o ./job-runner .


FROM python:3.13-alpine

# Didn't originally include build tools but eventually needed them due to luarocks compiling stuff with gcc
# so this may as well be the only container. Kept the separated builder above for now though.
RUN apk --no-cache add build-base ca-certificates libc6-compat wget lua5.3 lua5.3-dev luarocks
WORKDIR /run

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /build/job-runner /run

# Copy PIQUANT here
COPY ./Piquant3.2.17 /run/Piquant

# Set up a python venv so it's ready. We may not be running python all the time but at least we're not doing this each time
RUN python3.13 -m venv .

RUN wget https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem -O global-bundle.pem
RUN chmod +x ./job-runner

# Command to run the executable
ENTRYPOINT ["./job-runner"]
