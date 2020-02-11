package main

import (
	"../Cryptography/myRsa"
)

func main()  {
	/*  对称加密
	fmt.Println("des 加解密")
	key := []byte("1234abcd")  //密钥长度，8个字节
	src := []byte("创建一个密码分组模式的接口对象")
	cipherText := Sym_encrypt.DesEncrypto(src,key)
	plainText := Sym_encrypt.DesDecrypto(cipherText,key)
	fmt.Printf("密文是：%s\n",cipherText)
	fmt.Printf("明文是：%s\n",plainText)

	fmt.Println("aes 加解密")
	key1 := []byte("1234abcd1234abcd")  //密钥长度，8个字节
	src1 := []byte("创建一个密码分组模式的接口对象")
	cipherText1 := Sym_encrypt.AesEncrypto(src1,key1)
	plainText1 := Sym_encrypt.AesDecrypto(cipherText1,key1)
	fmt.Printf("密文是：%s\n",cipherText1)
	fmt.Printf("明文是：%s\n",plainText1)

	 */
/*
	//rsa加解密
	myRsa.GenerateRsaKey(2048)  //生成密钥对

	src := []byte("something just like this")  //需加密的字符串
	cipherText := myRsa.RSAEncrypto(src,"public.pem")
	plainText := myRsa.RSADecrypto(cipherText,"private.pem")
	fmt.Println(string(plainText))
*/
	myRsa.MyHash()
}
