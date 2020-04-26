module Mi_house/PostSession

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

replace Mi_house/PostSession/proto/PostSession => F:\Software\go_path\src\Mi_house\PostSession\proto\PostSession

replace Mi_house/PostSession/handler => F:\Software\go_path\src\Mi_house\PostSession\handler

go 1.13

require (
	Mi_house/PostSession/handler v0.0.0-00010101000000-000000000000
	Mi_house/PostSession/proto/PostSession v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
)
