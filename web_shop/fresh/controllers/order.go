package controllers

import (
	"../models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"github.com/smartwalle/alipay"
	"log"
	"strconv"
	"strings"
	"time"
)

type OrderController struct {
	beego.Controller
}

func(this *OrderController) ShowOrder(){
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

		id ,err:= strconv.Atoi(skuid)
		if err != nil {
			log.Println("商品id不合法", err)
			this.Redirect("/user/cart", 302)
			return
		}
		//查询商品数据
		var goodsSku models.GoodsSKU
		goodsSku.Id = id
		o.Read(&goodsSku)

		temp["goods"] = goodsSku
		//获取商品数量
		var count = 0
		source, _ := this.GetInt("source")
		//log.Println("本次订单显示的请求来源(1商品详情  0购物车)：", source)
		// 判断请求来源,若直接来自商品详情，则使用商品详情的购买数量数量，购物车不动
		if source == 1 {
			this.Data["source"] = 1
			count, err = this.GetInt("goodsCount")
			if err != nil {
				log.Println("商品数量错误", err)
				return
			}
			// 否则使用购物车的数量，并清空购物车对于商品
		} else {
			this.Data["source"] = 0
			count, err = redis.Int(conn.Do("hget", "cart_"+strconv.Itoa(user.Id), skuid))
			if err != nil {
				log.Println("商品数量错误", err)
				this.Redirect("/user/cart", 302)
				return
			}
		}

		//count,_ :=redis.Int(conn.Do("hget","cart_"+strconv.Itoa(user.Id),id))  //("hget",key值,value值)
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
	ids := skuid[1:len(skuid)-1]  //字符串切割，去掉[1,3,5]中两边的[]

	skuids := strings.Split(ids," ")  //再将获得的字符以空格隔开

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
		i := 3   //循环判断库存（可能下单成功但提交订单失败。下单的时候有库存，再提交时别人买了。重新拿到precount与数据库判断

		for i> 0{
			o.Read(&goods)

			var orderGoods models.OrderGoods

			orderGoods.GoodsSKU = &goods  //商品信息
			orderGoods.OrderInfo = &order  //订单信息
			//获取count
			count ,_ :=redis.Int(conn.Do("hget","cart_"+strconv.Itoa(user.Id),id))

			if count > goods.Stock{
				resp["code"] = 2
				resp["errmsg"] = "商品库存不足"
				this.Data["json"] = resp
				o.Rollback()  //标识事务的回滚
				return
			}

			preCount := goods.Stock  //原来的库存

			//time.Sleep(time.Second * 5)
			//beego.Info(preCount,user.Id)

			orderGoods.Count = count

			orderGoods.Price = count * goods.Price  //小计

			o.Insert(&orderGoods)

			goods.Stock -= count  //计算库存
			goods.Sales += count

			//更新数据库库存（高级更新，orm.Params指定更新字段）返回更新的数据数（更新了几条数据
			//Filter("Stock",preCount)  #若数据库里的数据与先前取的数据一致则更新，不一致（库存变了）则不更新
			updateCount,_:=o.QueryTable("GoodsSKU").Filter("Id",goods.Id).Filter("Stock",preCount).Update(orm.Params{"Stock":goods.Stock,"Sales":goods.Sales})
			if updateCount == 0{  //说明没有更新
				if i >0 {
					i -= 1
					continue
				}
				resp["code"] = 3
				resp["errmsg"] = "商品库存改变,订单提交失败"
				this.Data["json"] = resp
				o.Rollback()  //标识事务的回滚
				return
			}else{  //更新成功
				conn.Do("hdel","cart_"+strconv.Itoa(user.Id),goods.Id) //更新成功后删掉购物车数据
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

//处理支付
func(this*OrderController)HandlePay(){
	//var aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtPLfffeuL
	// cVVBAZmiQuA7BtFGv7GKG6mWP7P+r9/koOTsICX6PObhGZwSR1BYtJhgcdimRI3UBBxyR3P4Ay7egp
	// cconLuyxqZYNfohfVRL48MfIyS7cHDdNkjz2r70gOLfjYwchM6ttkzftME0k4QLJf/Y+qbSCiWvZ+9YR
	// FmHo9Iq8juKDbnYkYmhoq7LDUxwVh7k9JeYW20kTIJecfNutCWGOcAC01jFymbNglrne8cUWet+qgY2Wh
	// GwEK1+2r1lWu+0azsNPPF3i3vVPAH1F2yxz6njhU26zO7A6+sB5Ff4DiULh3UAH9yID6LKJNBVJTpKobwidhF
	// qk3ip5UqQIDAQAB" // 可选，支付宝提供给我们用于签名验证的公钥，通过支付宝管理后台获取

	var privateKey = "MIIEogIBAAKCAQEAq+o+CXBURufYmdQe8oqdYjC9cYQhmaldiwBtt4PBQ0s0Bm0L"+
	"Ro/lxyyQ3fiBy3mWhFNT/gxKEuTJmjEY1KdoNJGJ5m6Kt8zfYHSyD9GVyhPbvcBl"+
	"zL0qDTWxFZ4acBuxquhWZSs7fsXAK+JjQAFPgWgPl2UW+9NvL8gmBMWz6QQF6nGf"+
	"ixez81+/iceu7lslHkb9I/BlvS+In69eX5tA5NL8mWtDScjZSF3S2mtNWvmYPuRF"+
	"3u9irWJQdsiD7Lmyk5PynnFkGeQsMidcL/qA1dbyjSHGBB9tgBw9SO1z0iZ88sGA"+
	"eyC1XunG/coL6OpFE/5Rz/z1a799dPmgIjhAVQIDAQABAoIBAHLI7i8DL2JAfyEG"+
	"vrsWzkq3XGYk4QJmUSz0Vk2HzUyPM+due27buYIpQXjT6mXfgx5LPPONZRAEbo9j"+
	"ZL7oH+2plzeia0CZrDQM9q8VMzw/0iJ6Cf6P9znmtZxHekOMWi5Q/w6BtTHJ+6vR"+
	"do9P0LQItRBS0OvHmp9+tdPN5XgYvNvZG7PN2MoC1Ihn8Wh6aHmdCb/5U7zHRsSX"+
	"54NJ0pVAqh2jfLwM2pZrprLQfUeX9K6wzdElt8UDLsR/LaSTajMkowh3A40MjDdc"+
	"PEtBztnbxXF8qvuBpJGimrhssv4x47sIHYy/t5y4VhctHb9lRz/xrBGSF1YccqQE"+
	"S1W0LYECgYEA35GCTmukUJPorry1GdeGGQEKtO49IMB6uYuHXQTkVVdzJJeh358j"+
	"l37R5UIkuLxmR0+BGBCLHhxhVOxM585vL33RT5PilqUuJHx8xCKIEWmizvbkOkmq"+
	"/FwL0vVSaaocgPMZMN5cqU93wAQq7G4Ws5yN+OeawVa0YqFoSb/1UkUCgYEAxNqF"+
	"PaXEEqaGmdslUPoSq19e63/Cb2/M4UMtcgx2ygv7BbtKzA0oAj60WPg28qVaydK5"+
	"QhV27NSF6my3b05W+AHLQmALAW4cfGPFOhBKKSZtu+lBoGoIO7g5M7OdCD5z0VsZ"+
	"dxjkLNfbCVRsfRIiVWxLXHscE2Oja5sAQJjcHtECgYBi6HsfupuNUoKEm9P7RNB8"+
	"y2szNJsynDUvVT1jt8BUyUWVkMf6qOsUIPU3WhXo7ODK96+DN/50KdYdvhUgEMI7"+
	"3ov1itWaFCXf2ntQQn60RWnT6oJ6DxFoaEiK4iG9wZaurd4dMqMH/LJb7vwWcXX1"+
	"rHNbBQwAMR3I7OO4jQY5iQKBgDS8adeDEg0BCawsFkokRF4etZoXZupdpCBw07rq"+
	"9ePY4828SCR0HdTEIaJ4FJhajgDlb/XxNCPsJ1gO+fFaplfYezBz9F7UtdLpE4wj"+
	"kwqFHf6qVYd5AG2nQ8PDCX2M5ZAffKn4RCBzT12nfUzgT8xcr6+9116mYcBRDPdB"+
	"PC8RAoGAVipJ/4rO9iAXsdl6VwoBjXBE7Ag09GAN8bAXtglZ951Uyq5MsSHGZG+c"+
	"HVIaBad4EgtXrLdMn062IUPl3S9QK0kFFcGK5zZv/LIq/DEDWGA3upe9rIqXvzs2"+
	"IVCQ6SHtp3BM5kGz/fIDWBwreMqYZRSNumap6FxNzscDS1DoYk8=" // 必须，上一步中使用 RSA签名验签工具 生成的私钥

	var appId = "2016102000727997"
	var client,_ = alipay.New(appId, privateKey, false)

	//获取数据
	orderId := this.GetString("orderId")
	totalPrice := this.GetString("totalPrice")

	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://xxx"   //公网ip
	p.ReturnURL = "http://192.168.0.5:8080/user/payok"  //自己的ip
	p.Subject = "天天生鲜购物平台"  //主题
	p.OutTradeNo = orderId
	p.TotalAmount = totalPrice
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, err = client.TradePagePay(p)   //返回的URL地址作return用
	if err != nil {
		fmt.Println(err)
	}

	var payURL = url.String()
	this.Redirect(payURL,302)
}

//支付成功
func(this *OrderController) PayOk(){
	//获取数据
	//支付成功返回的字符串
	//http://192.168.0.5:8080/user/userCenterOrder?charset=utf-8&out_trade_no=20200208231138431&method=
	// alipay.trade.page.pay.return&total_amount=116.00&sign=MqV7pDNUYMLyIJRuhq60KT%2FR%2FR5gzl5TJJNRj%
	// 2FgFyUmBjRboJT5TU6ajyqx1P%2Fp46zW7OK2%2FrgVIFhj2FcHLEBanrsV%2F5jyaVNkANCl088X%2BcvQ2oGfQJAGDWckO
	// FrIoTHWzsH79vsulBCYTTgBcUL%2F%2FxO9nKktp1pKgNn1sVo%2BZFrcnDw%2FjUvZaCtUrx33TfVJfr5L7c6%2F5g9JZjz
	// yTMyFzE%2BpDemtxdtbAlX8ONhurY2BnFDeMZE2x1%2FsHmUP0KKS2f9R01FO1XKhC63hAhc5gRCfViF9zDYAmLZv%2BxmHk
	// kRfrJven33NFnaJ7VhUxDku9cHRM0qXWiZ%2FPPbUyRQ%3D%3D&trade_no=2020020922001490861000041877&
	// auth_app_id=2016102000727997&version=1.0&app_id=2016102000727997&sign_type=RSA2&seller_id=2088102180438940
	// &timestamp=2020-02-09+15%3A34%3A54
	orderId := this.GetString("out_trade_no")
	//log.Println("本次订单ID：", orderId)

	//校验数据
	if orderId ==""{
		beego.Info("支付返回数据错误")
		this.Redirect("/user/userCenterOrder",302)
		return
	}

	//操作数据
	o := orm.NewOrm()
	//_, err := o.QueryTable("OrderInfo").Filter("OrderId", orderId).Update(orm.Params{"Orderstatus": 0})
	//if err != nil {
	//	log.Println("更新订单数据失败")
	//	this.Redirect("/user/userCenterOrder", 302)
	//	return
	//}

	count,_:=o.QueryTable("OrderInfo").Filter("OrderId",orderId).Update(orm.Params{"Orderstatus":0})
	if count == 0{
		beego.Info("更新数据失败")
		this.Redirect("/",302)
		return
	}

	//返回视图
	this.Redirect("/user/userCenterOrder",302)
}

