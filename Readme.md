# Command to generate grpc code & protocol buffer messages
> protoc --proto_path=api --go_out=invoicer --go_opt=module=github.com/shamhub/go_grpc/invoicer --go-grpc_out=invoicer --go-grpc_opt=module=github.com/shamhub/go_grpc/invoicer invoicer.proto


# Set this require directive in go.mod
```
require (
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.33.0
)
```


# Generated files
> `invoicer.pb.go`, which contains all the protocol buffer code to populate, serialize, and retrieve request and response message types

> `invoicer_grpc.pb.go`, which contains an interface type (or stub) for clients to call with the methods defined in the Invoicer service, and an interface type for servers to implement, also with the methods defined in the Invoicer service