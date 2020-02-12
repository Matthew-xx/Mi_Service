package Mac

import (
	"crypto/hmac"
	"crypto/sha1"
)

//生成消息认证码
func GenerateHmac(plainText,key []byte) []byte {
	//创建哈希接口，需指定使用的哈希算法和密钥
	myhash := hmac.New(sha1.New,key)
	//给哈希对象添加数据
	myhash.Write(plainText)
	//计算散列值
	hashText := myhash.Sum(nil)
	//hex.EncodeToString(hashText)  //如果是网络通信，则最好编码
	return hashText
}

//验证消息认证码
func VerifyHmac(plainText,key,hashText []byte) bool {
	//创建哈希接口，需指定使用的哈希算法和密钥
	myhash := hmac.New(sha1.New,key)
	//给哈希对象添加数据
	myhash.Write(plainText)
	//计算散列值
	hashMac := myhash.Sum(nil)
	//两个散列值比较
	return hmac.Equal(hashText,hashMac)
}
