package handler

import (
	"Mi_house/ihomeWeb/models"
	"Mi_house/ihomeWeb/utils"
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"

	POSTHOUSES "Mi_house/PostHouses/proto/PostHouses"
)

type PostHouses struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PostHouses) CallPostHouses(ctx context.Context, req *POSTHOUSES.Request, rsp *POSTHOUSES.Response) error {
	beego.Info("发布房源请求 PostHouses /api/v1.0/user/houses")
	// 初始化回复
	rsp.Error = utils.RECODE_OK
	rsp.ErrMsg = utils.RecodeText(rsp.Error)
	// 获取请求参数
	sessionID := req.GetSessionID()
	houseInfoJson := req.GetHouseInfo()
	// 解析houseInfo
	houseInfo := make(map[string]interface{})
	json.Unmarshal(houseInfoJson, &houseInfo)
	beego.Info("用户发送房源信息表单：", houseInfo)
	// 读取redis链接配置
	redisConf := map[string]string{
		"key":      utils.G_server_name,
		"conn":     utils.G_redis_addr + ":" + utils.G_redis_port,
	}
	// 将map转换为json
	redisConfJSON, _ := json.Marshal(redisConf)
	// 开始链接redis
	bm, err := cache.NewCache("redis", string(redisConfJSON))
	if err != nil {
		beego.Info("缓存查询失败", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 验证sessionID，并得到id
	reply := bm.Get(sessionID + "user_id")
	if reply == nil {
		beego.Info("缓存查询结果为空")
		rsp.Error = utils.RECODE_NODATA
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return nil
	}
	id, err := redis.Int(reply, nil)
	if err != nil {
		beego.Info("缓存数据类型错误", err)
		rsp.Error = utils.RECODE_DATAERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 准备插入数据
	house := models.House{
		Title:   houseInfo["title"].(string),
		Address: houseInfo["address"].(string),
		Unit:    houseInfo["unit"].(string),
		Beds:    houseInfo["beds"].(string),
		// 	"title":"上奥世纪中心",
		// "price":"666",
		// "area_id":"5",
		// "address":"西三旗桥东建材城1号",
		// "room_count":"2",
		// "acreage":"60",
		// "unit":"2室1厅",
		// "capacity":"3",
		// "beds":"双人床2张",
		// "deposit":"200",
		// "min_days":"3",
		// "max_days":"0",
		// "facility":["1","2","3","7","12","14","16","17","18","21","22"]
	}
	user := models.User{Id: id}
	house.User = &user
	house.Price, _ = strconv.Atoi(houseInfo["price"].(string))
	// 构造自定义引用类型的数据，注意用new开辟内存
	area := new(models.Area)
	area.Id, _ = strconv.Atoi(houseInfo["area_id"].(string))
	house.Area = area
	house.Room_count, _ = strconv.Atoi(houseInfo["room_count"].(string))
	house.Acreage, _ = strconv.Atoi(houseInfo["acreage"].(string))
	house.Capacity, _ = strconv.Atoi(houseInfo["capacity"].(string))
	house.Deposit, _ = strconv.Atoi(houseInfo["deposit"].(string))
	house.Min_days, _ = strconv.Atoi(houseInfo["min_days"].(string))
	house.Max_days, _ = strconv.Atoi(houseInfo["max_days"].(string))
	// 构造model切片，接收json里面的切片元素。
	//facility 并不能直接接收，一个一个遍历json切片元素，存入model切片对应位置
	fa := []*models.Facility{}
	for _, v := range houseInfo["facility"].([]interface{}) {
		beego.Info("设施：", v, reflect.TypeOf(v))
		id, _ := strconv.Atoi(v.(string))
		tmp := models.Facility{Id: id}
		fa = append(fa, &tmp)
	}
	house.Facilities = fa
	o := orm.NewOrm()
	// 数据model构造完毕，开始写入
	_, err = o.Insert(&house)
	if err != nil {
		beego.Info("插入数据错误", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	// 上面的插入方法不能讲设施数据插入进入，进入house表并没有设施字段，而是有专门的房屋——设施关系表单独维护
	// 因为这是一个多对多的关系。故上面的插入方法会丢失设施信息。
	// 解决方法：多对多插入
	_, err = o.QueryM2M(&house, "Facilities").Add(house.Facilities)
	if err != nil {
		beego.Info("插入数据错误", err)
		rsp.Error = utils.RECODE_DBERR
		rsp.ErrMsg = utils.RecodeText(rsp.Error)
		return err
	}
	beego.Info("多对多插入自动创建的房屋记录id", house.Id)
	rsp.HousID = strconv.Itoa(house.Id)
	return nil
}
