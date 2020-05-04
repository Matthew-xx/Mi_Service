module Mi_house/PostHousesImage

replace Mi_house/PostHousesImage/handler => F:\Software\go_path\src\Mi_house\PostHousesImage\handler

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

replace Mi_house/PostHousesImage/proto/PostHousesImage => F:\Software\go_path\src\Mi_house\PostHousesImage\proto\PostHousesImage

go 1.13

require (
	Mi_house/PostHousesImage/handler v0.0.0-00010101000000-000000000000
	Mi_house/PostHousesImage/proto/PostHousesImage v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/protoc-gen-micro v1.0.0
)
