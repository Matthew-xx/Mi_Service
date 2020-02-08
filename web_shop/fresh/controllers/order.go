package controllers

import (
	"../models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"time"
)

type OrderController struct {
	beego.Controller
}

func(this*OrderController)ShowOrder(){
	//获取数据
	skuids :=this.GetStrings("skuid")  //获取切片（因为是多个商品
	//beego.Info(skuids)
	//校验数据
	if len(skuids) == 0{
		beego.Info("请求数据错误")
		this.Redirect("/user/cart",302)
		return
	}

	//处理数据
	o := orm.NewOrm()
	conn,_ := redis.Dial("tcp","192.168.99.100:6379")
	defer conn.Close()
	//获取用户数据
	var user models.User
	userName := this.GetSession("userName")
	user.Name = userName.(string)
	o.Read(&user,"Name")

	goodsBuffer := make([]map[string]interface{},len(skuids))

	totalPrice := 0
	totalCount := 0
	for index,skuid := range skuids{
		temp := make(map[string]interface{})

		id ,_ := strconv.Atoi(skuid)
		//查询商品数据
		var goodsSku models.GoodsSKU
		goodsSku.Id = id
		o.Read(&goodsSku)

		temp["goods"] = goodsSku
		//获取商品数量
		count,_ :=redis.Int(conn.Do("hget","cart_"+strconv.Itoa(user.Id),id))  //("hget",key值,value值)
		temp["count"] = count
		//计算小计(单个商品
		amount := goodsSku.Price * count
		temp["amount"] = amount

		//计算总金额和总件数
		totalCount += count
		totalPrice += amount

		goodsBuffer[index] = temp  //不同的index对应的商品信息,将临时容器放在buffer里
	}

	this.Data["goodsBuffer"] = goodsBuffer

	//获取地址数据
	var addrs []models.Address    //切片，以便循环
	//关联用户表并通过Id过滤
	o.QueryTable("Address").RelatedSel("User").Filter("User__Id",user.Id).All(&addrs)
	this.Data["addrs"] = addrs
	this.Data["userName"] = userName

	//传递总金额和总件数
	this.Data["totalPrice"] = totalPrice
	this.Data["totalCount"] = totalCount
	transferPrice := 10
	this.Data["transferPrice"] = transferPrice
	this.Data["realyPrice"] = totalPrice + transferPrice

	//传递所有商品的id
	this.Data["skuids"] = skuids

	//返回视图
	this.TplName = "place_order.html"
}

//添加订单
func(this *OrderController) AddOrder(){
	//获取数据
	addrid,_ :=this.GetInt("addrid")
	payId,_ :=this.GetInt("payId")
	skuid := this.GetString("skuids")  //获取到的是切片
	ids := skuid[1:len(skuid)-1]

	skuids := strings.Split(ids," ")

	//beego.Error(skuids)
	//totalPrice,_ := this.GetInt("totalPrice")
	totalCount,_ := this.GetInt("totalCount")
	transferPrice,_ :=this.GetInt("transferPrice")
	realyPrice,_:=this.GetInt("realyPrice")

	resp := make(map[string]interface{})
	defer this.ServeJSON()
	//校验数据
	if len(skuids) == 0{
		resp["code"] = 1
		resp["errmsg"] = "数据库链接错误"
		this.Data["json"] = resp
		return
	}
	//处理数据
	//向订单表中插入数据
	o := orm.NewOrm()

	o.Begin()//标识事务的开始

	userName := this.GetSession("userName")
	var user models.User
	user.Name = userName.(string)
	o.Read(&user,"Name")  //获取user.id

	var order models.OrderInfo
	order.OrderId = time.Now().Format("2006010215030405")+strconv.Itoa(user.Id)  //用时间设置订单id
	order.User = &user
	order.Orderstatus = 1
	order.PayMethod = payId
	order.TotalCount = totalCount
	order.TotalPrice = realyPrice
	order.TransitPrice = transferPrice
	//查询地址
	var addr models.Address
	addr.Id = addrid
	o.Read(&addr)

	order.Address = &addr

	//执行插入操作
	o.Insert(&order)


	//向订单商品表中插入数据
	conn,_ :=redis.Dial("tcp","192.168.99.100:6379")

	for _,skuid := range skuids{
		id,_ := strconv.Atoi(skuid)

		var goods models.GoodsSKU
		goods.Id = id
		i := 3

		for i> 0{
			o.Read(&goods)

			var orderGoods models.OrderGoods

			orderGoods.GoodsSKU = &goods
			orderGoods.OrderInfo = &order
			//获取count
			count ,_ :=redis.Int(conn.Do("hget","cart_"+strconv.Itoa(user.Id),id))

			if count > goods.Stock{
				resp["code"] = 2
				resp["errmsg"] = "商品库存不足"
				this.Data["json"] = resp
				o.Rollback()  //标识事务的回滚
				return
			}

			preCount := goods.Stock

			time.Sleep(time.Second * 5)
			//beego.Info(preCount,user.Id)

			orderGoods.Count = count

			orderGoods.Price = count * goods.Price  //小计

			o.Insert(&orderGoods)

			goods.Stock -= count
			goods.Sales += count

			updateCount,_:=o.QueryTable("GoodsSKU").Filter("Id",goods.Id).Filter("Stock",preCount).Update(orm.Params{"Stock":goods.Stock,"Sales":goods.Sales})
			if updateCount == 0{
				if i >0 {
					i -= 1
					continue
				}
				resp["code"] = 3
				resp["errmsg"] = "商品库存改变,订单提交失败"
				this.Data["json"] = resp
				o.Rollback()  //标识事务的回滚
				return
			}else{
				conn.Do("hdel","cart_"+strconv.Itoa(user.Id),goods.Id)
				break
			}
		}

	}

	//返回数据
	o.Commit()  //提交事务
	resp["code"] = 5
	resp["errmsg"] = "ok"
	this.Data["json"] = resp

}
