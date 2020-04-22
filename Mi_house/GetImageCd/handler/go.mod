module Mi_house/GetImageCd/handler

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/GetImageCd/proto/GetImageCd => F:\Software\go_path\src\Mi_house\GetImageCd\proto\GetImageCd

go 1.13

require (
	Mi_house/GetImageCd/proto/GetImageCd v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868 // indirect
	github.com/astaxie/beego v1.12.1
)
