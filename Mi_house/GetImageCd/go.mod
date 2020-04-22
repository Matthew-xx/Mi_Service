module Mi_house/GetImageCd

replace Mi_house/GetImageCd/proto/GetImageCd => F:\Software\go_path\src\Mi_house\GetImageCd\proto\GetImageCd

replace Mi_house/GetImageCd/handler => F:\Software\go_path\src\Mi_house\GetImageCd\handler

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

go 1.13

require (
	Mi_house/GetImageCd/handler v0.0.0-00010101000000-000000000000
	Mi_house/GetImageCd/proto/GetImageCd v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	golang.org/x/image v0.0.0-20200119044424-58c23975cae1 // indirect
)
