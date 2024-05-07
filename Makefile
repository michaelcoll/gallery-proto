rebuild-all: gen doc

prepare:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

.PHONY: gen
gen:
	buf generate contracts

lint:
	buf lint

.PHONY: doc
doc:
	protoc --doc_out=./doc --doc_opt=markdown,photo.md contracts/photo/v1/photo.proto
	protoc --doc_out=./doc --doc_opt=markdown,daemon.md contracts/daemon/v1/daemon.proto
