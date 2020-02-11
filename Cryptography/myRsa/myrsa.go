package myRsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

//生成rsa的密钥对，并保存到文件中
func GenerateRsaKey(keySize int)  {
	//生成私钥
	privateKey,_ := rsa.GenerateKey(rand.Reader,keySize)
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	//	组织pem.block
	block := pem.Block{
		Type:"rsa private key",   //字符串随意
		Bytes:derText,
	}
	//pem编码
	file,err :=os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file,&block)
	file.Close()

	//生成公钥
	//从私钥中取出公钥
	publicKey := privateKey.PublicKey
	//使用x509标准序列化
	dertream,err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//将返回的数据放入pem.block
	block1 := pem.Block{
		Type:"rsa public key",
		Bytes:dertream,
	}
	//pem编码
	file,err = os.Create("public.pem")
	pem.Encode(file,&block1)
	file.Close()
}

//加密函数(公钥加密
func RSAEncrypto(plainText []byte,fileName string) []byte {
	//读取公钥
	file,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileinfo,_ := file.Stat()
	buffer := make([]byte,fileinfo.Size())
	file.Read(buffer)  //将文件内容读取到容器中
	file.Close()

	//pem解码
	block,_ := pem.Decode(buffer)
	pubInterface,_ := x509.ParsePKIXPublicKey(block.Bytes)
	pubKey := pubInterface.(*rsa.PublicKey)  //断言，看是否符合公钥结构

	//使用公钥加密
	cipherText,err := rsa.EncryptPKCS1v15(rand.Reader,pubKey,plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}

//解密函数
func RSADecrypto(cipherText []byte,fileName string) []byte {
	//读取私钥
	file,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileinfo,_ := file.Stat()
	buffer := make([]byte,fileinfo.Size())
	file.Read(buffer)  //将文件内容读取到容器中
	file.Close()

	//pem解码
	block,_ := pem.Decode(buffer)
	privKey,_ := x509.ParsePKCS1PrivateKey(block.Bytes) //拿到私钥

	//使用私钥解密
	plainText,err := rsa.DecryptPKCS1v15(rand.Reader,privKey,cipherText)
	if err != nil {
		panic(err)
	}
	return plainText
}


