proto生成go文件

普通编译：protoc --go_out=paths=source_relative:. /per.proto

rpc插件编译：protoc --go_out=plugins=grpc,paths=source_relative:. hello.proto