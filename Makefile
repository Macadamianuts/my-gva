CONFIG_PROTO_FILES=$(shell find config -name *.proto)
GENERATOR_PROTO_FILES=$(shell find plugin/generator/config -name *.proto)

ifeq ($(PLUGIN),)
PLUGIN            = email
else
endif

.PHONY: config

config:
	protoc --proto_path=./config \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./config \
	       $(CONFIG_PROTO_FILES)

generator:
	protoc --proto_path=./plugin/generator/config \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:./plugin/generator/config \
		   $(GENERATOR_PROTO_FILES)