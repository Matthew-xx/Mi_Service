module Mi_house/DeleteSession

replace Mi_house/DeleteSession/handler => F:\Software\go_path\src\Mi_house\DeleteSession\handler

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/DeleteSession/proto/DeleteSession => F:\Software\go_path\src\Mi_house\DeleteSession\proto\DeleteSession

go 1.13

require (
	Mi_house/DeleteSession/handler v0.0.0-00010101000000-000000000000
	Mi_house/DeleteSession/proto/DeleteSession v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
)
