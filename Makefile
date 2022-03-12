.PHONY:protoc
protoc:
	--proto_path="api/protobuf/greeting" --go_out="." --go-grpc_out="." greeting.proto

.PHONY:tidy
go: 
	go mod tidy

.PHONY:useless
useless:
这是一个用mingw中的make 命令实现在终端代替重复性代码的功能即 make protoc 就相当于使用了protoc --proto_path="api/protobuf/greeting" --go_out="." --go-grpc_out="." greeting.proto