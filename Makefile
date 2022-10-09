rebuild-all: gen doc

.PHONY: gen
gen:
	buf generate contracts

lint:
	buf lint

.PHONY: doc
doc:
	protoc --doc_out=./doc --doc_opt=markdown,photo.md contracts/photo/v1/photo.proto
	protoc --doc_out=./doc --doc_opt=markdown,daemon.md contracts/daemon/v1/daemon.proto
