generate_grpc_code:
	protoc \
	--proto_path=api \
	--go_out=invoicer \
	--go_opt=module=github.com/shamhub/go_grpc/invoicer \
	--go-grpc_out=invoicer \
	--go-grpc_opt=module=github.com/shamhub/go_grpc/invoicer \
	invoicer.proto

