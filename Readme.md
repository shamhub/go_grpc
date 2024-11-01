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


# Reference:
https://www.youtube.com/watch?v=gbrPMv_GuQY


# Testing
1) Test with bloomrpc
2) Sample client code
```
package main

import (
 "context"
 pb "github.com/pascalallen/grpc-go/helloworld"
 "google.golang.org/grpc"
 "google.golang.org/grpc/credentials/insecure"
 "log"
 "time"
)

func main() {
 conn, err := grpc.Dial("localhost:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
 if err != nil {
  log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
 }
 defer conn.Close()
 c := invoicer.NewInvoicerClient(conn)

 ctx, cancel := context.WithTimeout(context.Background(), time.Second)
 defer cancel()

 r, err := c.Create(ctx, &invoicer.CreateRequest{})
 if err != nil {
  log.Fatalf("error calling function Create: %v", err)
 }

 log.Printf("Response from gRPC server's Create function: %s", r.GetMessage())
}
```