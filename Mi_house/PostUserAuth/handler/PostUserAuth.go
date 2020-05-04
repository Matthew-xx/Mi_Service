package handler

import (
	"Mi_house/ihomeWeb/models"
	"Mi_house/ihomeWeb/utils"
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"reflect"

	POSTUSERAUTH "Mi_house/PostUserAuth/proto/PostUserAuth"
)

type PostUserAuth struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PostUserAuth) CallPostUserAuth(ctx context.Context, req *POSTUSERAUTH.Request, rsp *POSTUSERAUTH.Response) error {
	beego.Info("发送进行实名认证请求 PostUserAuth /api/v1.0/user/auth")

	// 1初始化回复
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 2获取请求参数
	sessionID := req.GetSessionID()
	realName := req.GetRealName()
	IDcard := req.GetIDCard()

	// 读取redis链接配置
	redisConf := map[string]string{
		"key":      utils.G_server_name,
		"conn":     utils.G_redis_addr + ":" + utils.G_redis_port,
	}
	// 将map转换为json
	redisConfJSON, _ := json.Marshal(redisConf)
	// 3开始链接redis
	bm, err := cache.NewCache("redis", string(redisConfJSON))
	if err != nil {
		beego.Info("缓存查询失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 4验证sessionID，并得到id
	reply := bm.Get(sessionID + "user_id")
	if reply == nil {
		beego.Info("缓存查询结果为空")
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	beego.Info(reply, reflect.TypeOf(reply))
	id, err := redis.Int(reply, nil)
	if err != nil {
		beego.Info("缓存数据类型错误", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	//5通过用户id更新用户相关数据
	user := models.User{Id: id, Real_name: realName, Id_card: IDcard}
	o := orm.NewOrm()
	_, err = o.Update(&user, "Real_name", "Id_card")
	if err != nil {
		beego.Info("数据库更新失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("实名信息更新成功", user)
	return nil
}

