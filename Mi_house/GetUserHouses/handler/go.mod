module Mi_house/GetUserHouses/handler

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

replace Mi_house/GetUserHouses/proto/GetUserHouses => F:\Software\go_path\src\Mi_house\GetUserHouses\proto\GetUserHouses

go 1.13

require (
	Mi_house/GetUserHouses/proto/GetUserHouses v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
)
