package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

/*
rpc包提供了通过网络或其他I/O连接对一个对象的导出方法的访问。服务端注册一个对象，
使它作为一个服务被暴露，服务的名字是该对象的类型名。注册之后，
对象的导出方法就可以被远程访问。服务端可以注册多个不同类型的对象（服务），
但注册具有相同类型的多个对象是错误的。
 */

//对象
type Pandaner int

//方法
/* func (t *T) MethodName(argType T1, replyType *T2) error */
//函数关键字（对象）函数名（对端发送的内容，返回给对端的内容）错误值
func (this *Pandaner) Getinfo(argType int,replyType *int) error {
	fmt.Println("打印对端发送过来的内容：",argType)

	//修改内容
	*replyType = argType + 12306;

	return nil
}

func panda(w http.ResponseWriter ,r *http.Request)  {
	io.WriteString(w,"hello my love")
}

func main()  {
	/* 1、web操作*/

	//页面请求
	http.HandleFunc("/panda",panda)
	ln,err := net.Listen("tcp",":10086")
	if err != nil {
		fmt.Println("网络错误")
	}
	http.Serve(ln,nil)  //底层已经加了goroutine了，已经是高并发

	/*2、rpc操作 */

	//将类别实例化为对象
	pd := new(Pandaner)
	//服务器注册对象
	rpc.Register(pd)
	rpc.HandleHTTP()

}
