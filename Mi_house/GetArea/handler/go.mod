module Mi_house/GetArea/handler

go 1.13

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/GetArea/proto/GetArea => F:\Software\go_path\src\Mi_house\GetArea\proto\GetArea

require (
	Mi_house/GetArea/proto/GetArea v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/astaxie/beego v1.12.1
	github.com/julienschmidt/httprouter v1.2.0
	github.com/micro/go-micro v1.18.0
)
