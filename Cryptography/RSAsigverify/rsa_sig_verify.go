package RSAsigverify

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"os"
)

//RSA签名认证

//RSA签名（私钥
func SignatureRSA(plainText []byte,fileName string) []byte {
	//打开磁盘私钥文件
	file,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	//读取私钥文件中的内容
	fileinfo,err := file.Stat()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte,fileinfo.Size())
	file.Read(buffer)
	file.Close()

	//使用pem对数据进行解码，得到block结构体
	block ,_ := pem.Decode(buffer)
	//使用x509解析成私钥结构体，得到私钥
	privateKey,err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	//创建一个哈希对象
	//sha512.Sum512(plainText)
	myHash := sha512.New()
	myHash.Write(plainText)
	hashText := myHash.Sum(nil)

	//使用RSA对散列值签名
	sigText,err := rsa.SignPKCS1v15(rand.Reader,privateKey,crypto.SHA512,hashText)
	if err != nil {
		panic(err)
	}

	return sigText
}

//RSA签名验证
func VerifyRSA(plainText,sigText []byte,pubfileName string) bool {
	//打开磁盘公钥文件
	file,err := os.Open(pubfileName)
	if err != nil {
		panic(err)
	}

	//读取私钥文件中的内容
	fileinfo,err := file.Stat()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte,fileinfo.Size())  //文件长度来创建切片容器
	file.Read(buffer)
	file.Close()

	//使用pem对数据进行解码，得到block结构体
	block ,_ := pem.Decode(buffer)
	//对block.bytes解析得到接口
	pubinterface,_ := x509.ParsePKIXPublicKey(block.Bytes)
	//进行类型断言，得到公钥结构体
	publicKey := pubinterface.(*rsa.PublicKey)
	//对原始消息进行哈希运行（和签名时使用的算法一致）得到散列值
	hashText := sha512.Sum512(plainText)  //是一个定长的数组
	//签名认证
	err = rsa.VerifyPKCS1v15(publicKey,crypto.SHA512,hashText[:],sigText)
	if err == nil {
		return true
	}
	return false
}

