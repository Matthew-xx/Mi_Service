package Elliptic_curve

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
)

//椭圆曲线签名认证

//生成密钥对
func GenerateEccKey()  {
	//使用ecdsa生成密钥对
	privateKey,err := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	if err != nil {
		panic(err)
	}
	//将私钥写入磁盘
	//使用x509序列化
	derText,err := x509.MarshalECPrivateKey(privateKey)  //序列化为字符串
	if err != nil {
		panic(err)
	}
	//将得到的切片字符串放入Pem.block结构体中
	block := pem.Block{
		Type:"ecdsa private key",
		Bytes:derText,
	}

	//使用pem编码
	file,err := os.Create("eccPrivate.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file,&block)  //写入磁盘
	file.Close()

	//从私钥中读取公钥
	pubilcKey := privateKey.PublicKey
	//使用x509序列化
	derText1,err := x509.MarshalPKIXPublicKey(&pubilcKey)
	//将序列化结果放到pem.block块中
	block1 := pem.Block{
		Type:"ecdsa public key",
		Bytes:derText1,
	}
	//使用pem编码
	file1,err := os.Create("eccPublic.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file1,&block1)
	file1.Close()
}

//ecc签名（私钥
func EccSignature(plainText []byte,privName string) (rText,sText []byte) {
	//打开磁盘私钥文件
	file,err := os.Open(privName)
	if err != nil {
		panic(err)
	}

	//读取私钥文件中的内容
	fileinfo,err := file.Stat()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte,fileinfo.Size())  //切片容器长度
	file.Read(buffer)  //存储到切片容器
	file.Close()
	//使用pem对数据进行解码，得到block结构体
	block ,_ := pem.Decode(buffer)
	//使用x509对私钥进行还原得到私钥
	privateKey,err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对原始数据进行哈希运算得到是散列值
	hashText := sha1.Sum(plainText)
	//进行数字签名
	r,s,err := ecdsa.Sign(rand.Reader,privateKey,hashText[:]) //返回两个整型指针
	if err != nil {
		panic(err)
	}
	//对r,s内存中的数据进行格式化，转换成切片（将内存中的数据保存到字符串中
	rText,err = r.MarshalText()
	if err != nil {
		panic(err)
	}
	sText,err = s.MarshalText()
	if err != nil {
		panic(err)
	}

	return
}

//ecc签名认证
func EccVerify(plainText,rText,sText []byte,pubFile string) bool {
	file,err := os.Open(pubFile)
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
	pubInterface,_ := x509.ParsePKIXPublicKey(block.Bytes)
	//进行类型断言，得到公钥结构体
	publicKey := pubInterface.(*ecdsa.PublicKey)

	//对原始数据进行hash运算得到散列值（一个定长的数组
	hashText := sha1.Sum(plainText)
	//将rText,sText转换成int类型
	var r,s big.Int
	r.UnmarshalText(rText)
	s.UnmarshalText(sText)
	//签名认证
	bl := ecdsa.Verify(publicKey,hashText[:],&r,&s)

	return bl
}

