module Mi_house/GetUserInfo

replace Mi_house/GetUserInfo/handler => F:\Software\go_path\src\Mi_house\GetUserInfo\handler

replace Mi_house/GetUserInfo/proto/GetUserInfo => F:\Software\go_path\src\Mi_house\GetUserInfo\proto\GetUserInfo

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

go 1.13

require (
	Mi_house/GetUserInfo/handler v0.0.0-00010101000000-000000000000
	Mi_house/GetUserInfo/proto/GetUserInfo v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/astaxie/beego v1.12.1
	github.com/garyburd/redigo v1.6.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd // indirect
)
