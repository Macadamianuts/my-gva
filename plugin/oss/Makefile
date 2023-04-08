CONFIG_PROTO_FILES=$(shell find config -name *.proto)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./config \
	       --proto_path=../../third_party \
 	       --go_out=paths=source_relative:./config \
	       $(CONFIG_PROTO_FILES)