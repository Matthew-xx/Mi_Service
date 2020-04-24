module Mi_house/PostReg/handler

replace Mi_house/ihomeWeb/models => F:\Software\go_path\src\Mi_house\ihomeWeb\models

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/PostReg/proto/PostReg => F:\Software\go_path\src\Mi_house\PostReg\proto\PostReg

go 1.13

require (
	Mi_house/PostReg/proto/PostReg v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro v1.18.0
)
