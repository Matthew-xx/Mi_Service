package handler

import (
	GETAREA "Mi_house/GetArea/proto/GetArea"
	"Mi_house/ihomeWeb/models"
	"Mi_house/ihomeWeb/utils"
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro/util/log"
	"time"
)

type GetArea struct{}

/*
// Call is a single request handler called via client.Call or the generated client code
func (e *GetArea) Call(ctx context.Context, req *getArea.Request, rsp *getArea.Response) error {
	log.Log("Received GetArea.Call request")
	rsp.Errmsg = "Hello " + req.String()
	return nil
}
*/

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetArea) Stream(ctx context.Context, req *GETAREA.StreamingRequest, stream GETAREA.GetArea_StreamStream) error {
	log.Logf("Received GetArea.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&GETAREA.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetArea) PingPong(ctx context.Context, stream GETAREA.GetArea_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&GETAREA.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}


// 服务所属方法的具体定义
func (e *GetArea) Call(ctx context.Context, req *GETAREA.Request, rsp *GETAREA.Response) error {
	beego.Info("请求地区信息：GetArea api/v1.0/areas")
	// 初始化错误码，默认成功信息
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)
	/* 1从缓存中获取数据，有数据则直接返回给前端
	2没有数据，则去数据库查找数据
	3将数据存到缓存中
	4把数据发送给前端。
	*/

	// {"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}
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
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 缓存中获取数据，指定key为area_info(下面存，用作area查询
	areaValueJSON := bm.Get("area_info")
	if areaValueJSON != nil {
		// 获取到数据了,打包发送给前端
		beego.Info("缓存查询成功")
		// beego.Info("JSON:", string(areaValueJSON.([]byte)))
		// 查询到的数据仍是json，这里将其转回map,返回的是一组地址，需要使用切片接收
		areaMap := make([]map[string]interface{}, 0)
		json.Unmarshal(areaValueJSON.([]byte), &areaMap)
		beego.Info("缓存中的数据：")
		// 把map打包成rsp
		for index, area := range areaMap {
			beego.Info(index, area)
			tmp := GETAREA.Response_Areas{Aid: int32(area["aid"].(float64)), Aname: area["aname"].(string)}
			rsp.Data = append(rsp.Data, &tmp)
		}
		// 如果缓存有，则无需去数据库了
		return nil
	}


	beego.Info("缓存未找到，开始查询数据库")
	// 创建orm句柄
	o := orm.NewOrm()

	// 查什么？用什么接收
	areas := []models.Area{}  //接收
	num, err := o.QueryTable("area").All(&areas, "Id", "Name") //查询全部
	if err != nil {
		beego.Info("数据库查询失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return err
	}
	if num == 0 {
		beego.Info("没有查询到数据")
		rsp.Error = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	// 将查询到的数据，存入缓存中
	areaJSON, err := json.Marshal(areas)
	//Put(key string, val interface{}, timeout time.Duration) error
	err = bm.Put("area_info", areaJSON, time.Second*time.Duration(utils.G_redis_expire))
	if err != nil {
		beego.Info("数据库缓存失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
	}


	// 将查询到的数据，按照proto的格式发送给web(前端
	for key, value := range areas {
		beego.Info(key, value)
		tmp := GETAREA.Response_Areas{Aid: int32(value.Id), Aname: value.Name}  //proto的格式
		rsp.Data = append(rsp.Data, &tmp)
	}
	return nil
}