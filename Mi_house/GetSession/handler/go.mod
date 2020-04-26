module Mi_house/GetSession/handler

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/GetSession/proto/GetSession => F:\Software\go_path\src\Mi_house\GetSession\proto\GetSession

go 1.13

require (
	Mi_house/GetSession/proto/GetSession v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/astaxie/beego v1.12.1
	github.com/garyburd/redigo v1.6.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/micro/go-micro v1.18.0
)
