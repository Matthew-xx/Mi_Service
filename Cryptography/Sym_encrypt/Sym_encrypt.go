package Sym_encrypt
//对称加密
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

//des的CBC加密
//填充函数，最后一个分组的字节数不够时填充。字节数刚好合适则添加一个新的分组
//填充的字节的值 == 缺少的字节的数

func paddingLastGroup(plainText []byte,blockSize int) []byte {
	//求出最后一个组中剩余的字节数 28%8 =3...4
	padNum := blockSize - len(plainText) % blockSize  //需填充的字节数
	//创建新的切片长度为padNum  每个值为byte(padNum)
	char := []byte{byte(padNum)}  //转换类型
	//填充部分
	newPlain := bytes.Repeat(char,padNum)  //对char重复n次
	//添加到原始明文中
	newPlainText := append(plainText,newPlain...)


	return newPlainText
}

//去掉填充的数据
func unPaddingLastGroup(plainText []byte) []byte {
	//将切片中最后一个字节取出
	length := len(plainText)
	lastChar := plainText[length-1]  //最后一个字节
	number := int(lastChar)

	return plainText[:length-number]
}

//des加密,plainText指明文
func DesEncrypto(plainText,key []byte) []byte {
	//建立底层使用des的密码接口
	block,err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//明文填充
	newText := paddingLastGroup(plainText,block.BlockSize())
	//对称加密，创建一个使用cbc分组接口
	iv := []byte("12345678")  //初始化向量
	blockMode := cipher.NewCBCEncrypter(block,iv)
	//加密
	cipherText := make([]byte,len(newText))   //容器用来存储密文
	blockMode.CryptBlocks(cipherText,newText)  //对第二个参数进行加密，加密的结果存储到第一个参数中

	return cipherText
}

//des解密
func DesDecrypto(cipherText,key []byte) []byte {
	//建立底层使用des的密码接口
	block,err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//创建一个使用cdc解密的接口
	iv := []byte("12345678")  //加解密的初始化向量需一致
	blockMode := cipher.NewCBCDecrypter(block,iv)
	//解密
	blockMode.CryptBlocks(cipherText,cipherText)  //第一参数是解密后得到的明文(覆盖存储），第二参数是密文
	//将解密后的明文删除其尾部加密时填充数据
	plainText := unPaddingLastGroup(cipherText)

	return plainText
}

//aes加密,分组模式ctr
func AesEncrypto(plainText,key []byte) []byte {
	//建立底层使用aes的密码接口
	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//对称加密，创建一个使用ctr分组接口
	iv := []byte("1234567812345678")  //初始化向量
	stream := cipher.NewCTR(block,iv)  //iv可理解为随机数种子。明文分组多长，iv的长度就多长
	//加密
	cipherText := make([]byte,len(plainText))   //容器用来存储密文
	stream.XORKeyStream(cipherText,plainText) //对第二个参数进行加密，加密的结果存储到第一个参数中

	return cipherText
}

//aes解密,按位异或操作，操作一次是加密，两次是解密。所以代码一致
func AesDecrypto(cipherText,key []byte) []byte {
	//建立底层使用des的密码接口
	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//创建一个使用ctr计数器解密的接口
	iv := []byte("1234567812345678")  //加解密的初始化向量需一致
	stream := cipher.NewCTR(block,iv)
	//解密
	stream.XORKeyStream(cipherText,cipherText)  //第一参数是解密后得到的明文(覆盖存储），第二参数是密文


	return cipherText
}