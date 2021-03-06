package handler

import (
	"Mi_house/ihomeWeb/models"
	"Mi_house/ihomeWeb/utils"
	"context"
	"encoding/json"
	"reflect"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	GETUSERINFO "Mi_house/GetUserInfo/proto/GetUserInfo"
)

type GetUserInfo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetUserInfo) CallGetUserInfo(ctx context.Context, req *GETUSERINFO.Request, rsp *GETUSERINFO.Response) error {
	beego.Info("获取用户信息 GetUserInfo /api/v1.0/users")
	// 初始化rsp
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 取得req参数
	sessionID := req.GetSessionID()
	//利用sessionID在缓存中读出用户ID
	redisConf := map[string]string{
		"key":      utils.G_server_name,
		"conn":     utils.G_redis_addr + ":" + utils.G_redis_port,
	}
	// 将map转换为json
	redisConfJSON, _ := json.Marshal(redisConf)

	// 链接redis,读出id
	bm, err := cache.NewCache("redis", string(redisConfJSON))
	if err != nil {
		beego.Info("缓存查询失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	//拿到key
	reply := bm.Get(sessionID + "user_id")
	if reply == nil {
		beego.Info("缓存查询结果为空")
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	beego.Info(reply, reflect.TypeOf(reply))  //打印下类型看是否一致
	//通过key拿到userid
	id, err := redis.Int(reply, nil)  //转换类型
	if err != nil {
		beego.Info("缓存数据类型错误", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 链接数据库，读出用户其他信息（通过id查询
	user := models.User{Id: id}
	o := orm.NewOrm()
	o.Read(&user)
	beego.Info(user)
	// rsp
	rsp.UserID = int64(user.Id)
	rsp.Name = user.Name
	rsp.Mobile = user.Mobile
	rsp.RealName = user.Real_name
	rsp.IDCard = user.Id_card
	rsp.AvatarUrl = user.Avatar_url

	return nil
}
