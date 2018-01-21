auth:
	grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./node --grpc_out=./node --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` auth.proto
	protoc -I . --go_out=plugins=grpc:golang auth.proto