# Command to generate grpc code & protocol buffer messages
> protoc --proto_path=api --go_out=invoicer --go_opt=module=github.com/shamhub/go_grpc/invoicer --go-grpc_out=invoicer --go-grpc_opt=module=github.com/shamhub/go_grpc/invoicer invoicer.proto


# Set this require directive in go.mod
```
require (
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.33.0
)
```