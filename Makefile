.PHONY:protoc
protoc:
	--protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/protobuf/greeting/greeting.proto

.PHONY:tidy
go: 
	go mod tidy

.PHONY:treer
treer:
	--treer -i "/node_modules|dist_electron|devtools|nouse|.git|.vscode/" -e "tree.md"

# .PHONY:useless
# useless:
# 这是一个用mingw中的make 命令实现在终端代替重复性代码的功能即 make protoc 就相当于使用了protoc --proto_path="api/protobuf/greeting" --go_out="." --go-grpc_out="." greeting.proto