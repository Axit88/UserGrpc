create:
	protoc --proto_path=proto proto/*.proto --go_out=storage-proto/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=storage-proto/

clean:
	rm storage-proto/*.go