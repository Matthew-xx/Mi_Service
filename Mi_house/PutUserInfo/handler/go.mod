module Mi_house/PutUserInfo/handler

replace Mi_house/PutUserInfo/proto/PutUserInfo => F:\Software\go_path\src\Mi_house\PutUserInfo\proto\PutUserInfo

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

go 1.13

require (
	Mi_house/PutUserInfo/proto/PutUserInfo v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
)
