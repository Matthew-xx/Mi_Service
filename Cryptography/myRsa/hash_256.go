package myRsa

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func MyHash()  {
	//简便使用（数据量少的情况下
	//sha256.Sum256([]byte("hello"))

	//较复杂情况下
	//创建hash接口对象
	myHash := sha256.New()
	//添加数据
	src := []byte("something just like this")
	myHash.Write(src)
	myHash.Write(src)
	myHash.Write(src)
	//三次拼接字符串（在大数据时，循环分开读取文件，每次都一个数据块，读一部分拼接一部分
	//计算结果
	res := myHash.Sum(nil)
	//格式化为16进制形式
	myStr := hex.EncodeToString(res)
	fmt.Printf("hash值:%s\n",myStr)
}
