# build environment for Golang

# build stage
FROM golang:1.14.6 AS build-env
WORKDIR /app

# Enable module mode
ENV GO111MODULE=on

# install protoc-gen-go
RUN go get github.com/golang/protobuf/protoc-gen-go
ENV PATH="$PATH:$(go env GOPATH)/bin"

# install protoc
ENV PB_REL="https://github.com/protocolbuffers/protobuf/releases"
RUN curl -LO $PB_REL/download/v3.12.1/protoc-3.12.1-linux-x86_64.zip
RUN apt-get update -y && \
    apt-get install -y unzip && \
    unzip protoc-3.12.1-linux-x86_64.zip && \
    cp bin/protoc /bin/ && \
    rm -rf *
COPY proto/ proto/
# generate gRPC package
RUN protoc --go_out=plugins=grpc:. proto/*.proto

COPY go.mod .
COPY go.sum .
# download Go modules
RUN go mod download

COPY . .

# build up Go APP
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go


# # production stage
FROM scratch
COPY --from=build-env /app/go /go
ENTRYPOINT ["/go"]