FROM golang:1.24-alpine AS builder

RUN apk add --no-cache build-base

COPY ./runner /build
ENV CGO_ENABLED=1

RUN cd /build && GOOS=linux GOARCH=amd64 go build -o ./job-runner .


FROM python:3.13-alpine

RUN apk --no-cache add build-base ca-certificates libc6-compat wget lua5.3 lua5.3-dev luarocks
WORKDIR /root

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /build/job-runner ./

COPY ./Piquant3.2.17 /build/Piquant

RUN wget https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem -O global-bundle.pem
RUN chmod +x ./job-runner

# RUN luarocks-5.3 install luafilesystem

# Command to run the executable
ENTRYPOINT ["./job-runner"]
