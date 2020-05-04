module Mi_house/GetUserHouses

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

replace Mi_house/GetUserHouses/proto/GetUserHouses => F:\Software\go_path\src\Mi_house\GetUserHouses\proto\GetUserHouses

replace Mi_house/GetUserHouses/handler => F:\Software\go_path\src\Mi_house\GetUserHouses\handler

go 1.13

require (
	Mi_house/GetUserHouses/handler v0.0.0-00010101000000-000000000000
	Mi_house/GetUserHouses/proto/GetUserHouses v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
)
