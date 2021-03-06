package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/tedcy/fdfs_client"
)

/* 将url加上 http://IP:PROT/ 前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求
func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url
	return domain_url
}

// md5工具函数
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 通过二进制流上传图片（参数为一个二进制和一个后缀
func UploadByBuffer(fileBuffer []byte, fileExt string) (string, error) {
	//绝对路径
	client, err := fdfs_client.NewClientWithConfig("f:/Software/go_path/src/github.com/tedcy/fdfs_client/fdfs.conf")
	defer client.Destory()
	if err != nil {
		//fmt.Println(err)
		beego.Info("fdfs客户端创建失败", err)
		return "", err
	}
	str, err := client.UploadByBuffer(fileBuffer, fileExt)
	if err != nil {
		beego.Info("fdfs客户端上传失败", err)
		return "", err
	}
	beego.Info("上传成功：", str)
	return str, nil

}
