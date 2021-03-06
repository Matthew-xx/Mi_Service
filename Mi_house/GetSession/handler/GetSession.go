package handler

import (
	"Mi_house/ihomeWeb/utils"
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	GETSESSION "Mi_house/GetSession/proto/GetSession"
)

type GetSession struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetSession) CallGetSession(ctx context.Context, req *GETSESSION.Request, rsp *GETSESSION.Response) error {
	beego.Info("获取登录状态 GetISession api/v1.0/session")
	sessionID := req.GetSessionID()
	// 初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 缓存中读Name
	redisConf := map[string]string{
		"key":      utils.G_server_name,
		"conn":     utils.G_redis_addr + ":" + utils.G_redis_port,
	}
	// 将map转换为json
	redisConfJSON, _ := json.Marshal(redisConf)

	// 链接redis
	bm, err := cache.NewCache("redis", string(redisConfJSON))
	if err != nil {
		beego.Info("缓存查询失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	//查询 "user_name" 与注册中的字段一致
	reply := bm.Get(sessionID + "user_name")
	if reply == nil {
		beego.Info("缓存查询结果为空")
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	str, _ := redis.String(reply, nil)
	rsp.Name = str

	return nil
}