# Grpc-API
Live Hub Server implementation for DOPI in Map

### Setup

#### 0. Install Go:

#### 1. Install Protocol buffer compiler (= `protoc`)

#### 2. Check `protoc` version:
    > protoc --version

Please ensure compiler version is 3+ or the proto WILL NOT compile.

#### 3. Install the `protoc` plugins for Go:
    > go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31
    > go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

#### 4. Update your `PATH` so that the `protoc` compiler can find the plugins:
    > export PATH="$PATH:$(go env GOPATH)/bin"

#### 5. Compile proto file and generate gRPC server files:
navigate to api/proto/server and excute the following command:

    > protoc --go_out=. --go-grpc_out=. .\server.proto

#### 6. update dependecies:
    > go mod tidy
Updates `go.mod` file with missing or unnecessary dependencies and creates a `go.sum` file.

#### 6. Run server:
    > go run main.go