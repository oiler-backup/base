PROTO_DIR := proto
MODULE_PATH := github.com/oiler-backup/base
PROTOC := protoc
GO_OUT := --go_out=.
GO_GRPC_OUT := --go-grpc_out=.
GO_OPT := --go_opt=module=$(MODULE_PATH)
GO_GRPC_OPT := --go-grpc_opt=module=$(MODULE_PATH)

all: generate-all

generate-all:
	@echo "Generating Go code for all .proto files in $(PROTO_DIR)..."
	$(PROTOC) \
    	$(GO_OUT) \
    	$(GO_GRPC_OUT) \
    	$(GO_OPT) \
    	$(GO_GRPC_OPT) \
    	$(PROTO_DIR)/*.proto
	@echo "Code generation completed."

generate: $(PROTO_DIR)/%.proto
	@echo "Generating Go code from $<..."
	$(PROTOC) \
    	$(GO_OUT) \
    	$(GO_GRPC_OUT) \
    	$(GO_OPT) \
    	$(GO_GRPC_OPT) \
    	$<
	@echo "Code generation for $< completed."

clean:
	@echo "Cleaning generated files..."
	rm -f $(PROTO_DIR)/*.pb.go $(PROTO_DIR)/*.grpc.pb.go
	@echo "Cleanup completed."

check-deps:
	@echo "Checking dependencies..."
	@which $(PROTOC) > /dev/null || (echo "Error: protoc is not installed. Please install it." && exit 1)
	@which protoc-gen-go > /dev/null || (echo "Error: protoc-gen-go is not installed. Run 'go install google.golang.org/protobuf/cmd/protoc-gen-go@latest'" && exit 1)
	@which protoc-gen-go-grpc > /dev/null || (echo "Error: protoc-gen-go-grpc is not installed. Run 'go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest'" && exit 1)
	@echo "All dependencies are installed."

help:
	@echo "Available targets:"
	@echo "  all           - Generate Go code for all .proto files in $(PROTO_DIR) (default target)."
	@echo "  generate-all  - Generate Go code for all .proto files in $(PROTO_DIR)."
	@echo "  generate      - Generate Go code for a specific .proto file (e.g., make generate PROTO_FILE=proto/backup.proto)."
	@echo "  clean         - Remove generated .pb.go and .grpc.pb.go files."
	@echo "  check-deps    - Check if required tools (protoc, protoc-gen-go, protoc-gen-go-grpc) are installed."
	@echo "  help          - Show this help message."

.PHONY: all generate-all generate clean check-deps help