module Mi_house/ihomeWeb/handler

go 1.13

replace Mi_house/ihomeWeb/models => ../models

replace Mi_house/ihomeWeb/utils => F:\Software\go_path\src\Mi_house\ihomeWeb\utils

replace Mi_house/GetArea/proto/GetArea => F:\Software\go_path\src\Mi_house\GetArea\proto\GetArea

replace Mi_house/GetImageCd/proto/GetImageCd => F:\Software\go_path\src\Mi_house\GetImageCd\proto\GetImageCd

replace Mi_house/GetIndex/proto/GetIndex => F:\Software\go_path\src\Mi_house\GetIndex\proto\GetIndex

replace Mi_house/GetSession/proto/GetSession => F:\Software\go_path\src\Mi_house\GetSession\proto\GetSession

replace Mi_house/PostSession/proto/PostSession => F:\Software\go_path\src\Mi_house\PostSession\proto\PostSession

replace Mi_house/DeleteSession/proto/DeleteSession => F:\Software\go_path\src\Mi_house\DeleteSession\proto\DeleteSession

replace Mi_house/GetSmsCd/proto/GetSmsCd => F:\Software\go_path\src\Mi_house\GetSmsCd\proto\GetSmsCd

replace Mi_house/PostReg/proto/PostReg => F:\Software\go_path\src\Mi_house\PostReg\proto\PostReg

replace Mi_house/GetUserInfo/proto/GetUserInfo => F:\Software\go_path\src\Mi_house\GetUserInfo\proto\GetUserInfo

replace Mi_house/PostAvatar/proto/PostAvatar => F:\Software\go_path\src\Mi_house\PostAvatar\proto\PostAvatar

replace Mi_house/PostUserAuth/proto/PostUserAuth => F:\Software\go_path\src\Mi_house\PostUserAuth\proto\PostUserAuth

replace Mi_house/GetUserHouses/proto/GetUserHouses => F:\Software\go_path\src\Mi_house\GetUserHouses\proto\GetUserHouses

replace Mi_house/PostHouses/proto/PostHouses => F:\Software\go_path\src\Mi_house\PostHouses\proto\PostHouses

require (
	Mi_house/DeleteSession/proto/DeleteSession v0.0.0-00010101000000-000000000000
	Mi_house/GetArea/proto/GetArea v0.0.0-00010101000000-000000000000
	Mi_house/GetImageCd/proto/GetImageCd v0.0.0-00010101000000-000000000000
	Mi_house/GetIndex/proto/GetIndex v0.0.0-00010101000000-000000000000
	Mi_house/GetSession/proto/GetSession v0.0.0-00010101000000-000000000000
	Mi_house/GetSmsCd/proto/GetSmsCd v0.0.0-00010101000000-000000000000
	Mi_house/GetUserHouses/proto/GetUserHouses v0.0.0-00010101000000-000000000000
	Mi_house/GetUserInfo/proto/GetUserInfo v0.0.0-00010101000000-000000000000
	Mi_house/PostAvatar/proto/PostAvatar v0.0.0-00010101000000-000000000000
	Mi_house/PostHouses/proto/PostHouses v0.0.0-00010101000000-000000000000
	Mi_house/PostReg/proto/PostReg v0.0.0-00010101000000-000000000000
	Mi_house/PostSession/proto/PostSession v0.0.0-00010101000000-000000000000
	Mi_house/PostUserAuth/proto/PostUserAuth v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/models v0.0.0-00010101000000-000000000000
	Mi_house/ihomeWeb/utils v0.0.0-00010101000000-000000000000
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868 // indirect
	github.com/astaxie/beego v1.12.1
	github.com/julienschmidt/httprouter v1.2.0
	github.com/micro/go-micro v1.18.0
)
