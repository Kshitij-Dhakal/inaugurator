CREATE_PROTO_PATH= pkg/create/
CREATE_PROTO_FILE= $(CREATE_PROTO_PATH)*.proto
CREATE_PROTO_DEST= ../../../

dev:
	export INAUGURATOR_ENV=dev

proto: 
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc -I=$(CREATE_PROTO_PATH) --go_out=$(CREATE_PROTO_DEST) $(CREATE_PROTO_FILE)
