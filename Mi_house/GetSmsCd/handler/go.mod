module Mi_house/GetSmsCd/handler

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/GetSmsCd/proto/GetSmsCd => F:\Software\go_path\src\Mi_house\GetSmsCd\proto\GetSmsCd

go 1.13

require (
	Mi_house/GetSmsCd/proto/GetSmsCd v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190808125512-07798873deee
	github.com/astaxie/beego v1.12.1
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/micro/go-micro v1.18.0
)
