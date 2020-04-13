package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	webTest "myTest"
)
//(传出，传入)
func WebTestCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("前端发送过来的json转换成map打印：",request["name"])

	// call the backend service
	webTestClient := webTest.NewMyTestService("go.micro.srv.myTest", client.DefaultClient)  //go.micro.srv.myTest是要请求的服务名
	rsp, err := webTestClient.Call(context.TODO(), &webTest.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("服务端做完处理回传的信息:",rsp.Msg)

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
