module Mi_house/GetHouseInfo/handler

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

replace Mi_house/GetHouseInfo/proto/GetHouseInfo => F:\Software\go_path\src\Mi_house\GetHouseInfo\proto\GetHouseInfo

go 1.13

require (
	Mi_house/GetHouseInfo/proto/GetHouseInfo v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
)
