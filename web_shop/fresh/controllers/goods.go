package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"../models"
	"github.com/garyburd/redigo/redis"
	"math"
	"strconv"
)

type GoodsController struct {
	beego.Controller
}

//使用beego控制器，父类，子类的所有控制器都能调用方法
func GetUser(this *beego.Controller) string{
	userName := this.GetSession("userName")
	if userName == nil{
		this.Data["userName"] = ""
	}else{
		this.Data["userName"] = userName.(string)  //显示用户名
		return userName.(string)
	}
	return ""
}


func PageTool(pageCount int,pageIndex int)[]int{

	var pages []int
	if pageCount <= 5{
		pages = make([]int,pageCount)
		for i,_ := range pages{
			pages[i] = i + 1
		}
		//pages = [1,2,..,pageCount]
	}else if pageIndex <= 3{
		//pages := make([]int,5)
		pages = []int{1,2,3,4,5}
	}else if pageIndex > pageCount - 3 {
		//pages = [6, 7, 8, 9, 10] 。显示当前页码和前面2页+后面3页
		pages = []int{pageCount -4,pageCount - 3,pageCount - 2,pageCount -1 ,pageCount}  //最后5页（i为9,10时
	}else {
		pages = []int{pageIndex - 2,pageIndex -1 ,pageIndex,pageIndex + 1, pageIndex + 2}
	}
	return pages

}


//展示首页
func(this *GoodsController) ShowIndex(){
	GetUser(&this.Controller)
	o := orm.NewOrm()
	//获取类型数据
	var goodsTypes []models.GoodsType
	o.QueryTable("GoodsType").All(&goodsTypes)
	this.Data["goodsTypes"] = goodsTypes

	//获取轮播图数据
	var indexGoodsBanner []models.IndexGoodsBanner
	o.QueryTable("IndexGoodsBanner").OrderBy("Index").All(&indexGoodsBanner)
	this.Data["indexGoodsBanner"] = indexGoodsBanner

	//获取促销商品数据
	var promotionGoods []models.IndexPromotionBanner
	o.QueryTable("IndexPromotionBanner").OrderBy("Index").All(&promotionGoods)
	this.Data["promotionsGoods"] = promotionGoods

	//首页展示商品数据
	goods := make([]map[string]interface{},len(goodsTypes)) //有多少个类型就多少个切片（先类型再商品

	//大类型的map，值是interface
	//向切片interface中插入类型数据
	for index, value := range goodsTypes{
		//获取对应类型的首页展示商品
		temp := make(map[string]interface{})
		temp["type"] = value
		goods[index] = temp
	}
	//商品数据
	for _,value := range goods{
		var textGoods []models.IndexTypeGoodsBanner
		var imgGoods []models.IndexTypeGoodsBanner
		//获取文字商品数据（RelatedSel关联两个表
		o.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsType","GoodsSKU").OrderBy("Index").Filter("GoodsType",value["type"]).Filter("DisplayType",0).All(&textGoods)
		//获取图片商品数据
		o.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsType","GoodsSKU").OrderBy("Index").Filter("GoodsType",value["type"]).Filter("DisplayType",1).All(&imgGoods)

		value["textGoods"] = textGoods
		value["imgGoods"] = imgGoods
	}
	this.Data["goods"] = goods
	cartCount := GetCartCount(&this.Controller)
	this.Data["cartCount"] = cartCount



	this.TplName = "index.html"
}

func ShowLaout(this*beego.Controller){
	//查询类型
	o := orm.NewOrm()
	var types []models.GoodsType
	o.QueryTable("GoodsType").All(&types)
	this.Data["types"] = types
	//获取用户信息
	GetUser(this)
	//指定layout
	this.Layout = "goodsLayout.html"
}

//展示商品详情
func(this *GoodsController) ShowGoodsDetail(){
	//获取数据
	id,err := this.GetInt("id")
	//校验数据
	if err != nil{
		beego.Error("浏览器请求错误")
		this.Redirect("/",302)
		return
	}
	//处理数据
	o := orm.NewOrm()
	var goodsSku models.GoodsSKU
	goodsSku.Id = id
	//o.Read(&goodsSku)
	o.QueryTable("GoodsSKU").RelatedSel("GoodsType","Goods").Filter("Id",id).One(&goodsSku)

	//获取同类型时间靠前的两条商品数据
	var goodsNew []models.GoodsSKU
	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType",goodsSku.GoodsType).OrderBy("Time").Limit(2,0).All(&goodsNew)
	this.Data["goodsNew"] = goodsNew

	//返回视图
	this.Data["goodsSku"] = goodsSku

	//添加历史浏览记录(记录什么时候添加：登录之后查看详情添加、获取：用户中心获取、存储：Redis存储）
	//判断用户是否登录
	userName := this.GetSession("userName")
	if userName != nil{
		//查询用户信息
		o := orm.NewOrm()
		var user models.User
		user.Name = userName.(string)  //赋值，断言string类型
		o.Read(&user,"Name")  //Name字段查询
		//添加历史记录,用redis存储
		conn,err := redis.Dial("tcp","192.168.99.100:6379")
		defer conn.Close()
		if err != nil{
			beego.Info("redis链接错误")
		}
		//把以前相同商品的历史浏览记录删除
		conn.Do("lrem","history_"+strconv.Itoa(user.Id),0,id) // 0表示全删，<0表从右往左删n个值
		//添加新的商品浏览记录
		conn.Do("lpush","history_"+strconv.Itoa(user.Id),id)  //("lpush",key值,商品Id)


	}

	ShowLaout(&this.Controller)
	cartCount := GetCartCount(&this.Controller)
	this.Data["cartCount"] = cartCount  //购物车数量
	this.TplName = "detail.html"

}

//展示商品列表页
func(this *GoodsController) ShowList(){
	//获取数据
	id,err := this.GetInt("typeId")
	//校验数据
	if err != nil{
		beego.Info("请求路径错误")
		this.Redirect("/",302)
		return
	}
	//处理数据
	ShowLaout(&this.Controller)  //类型
	//获取新品
	o := orm.NewOrm()
	var goodsNew []models.GoodsSKU
	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",id).OrderBy("Time").Limit(2,0).All(&goodsNew) //从0开始拿2条数据
	this.Data["goodsNew"] = goodsNew

	//获取商品
	var goods []models.GoodsSKU
	//o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",id).All(&goods)

	//分页实现
	//获取pageCount
	count,_ := o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",id).Count()
	pageSize := 2
	pageCount := math.Ceil(float64(count)/float64(pageSize))  //总页码

	pageIndex,err := this.GetInt("pageIndex")  //当前页码
	if err != nil{
		pageIndex = 1
	}
	pages := PageTool(int(pageCount),pageIndex)  //显示页码
	this.Data["pages"] = pages
	this.Data["typeId"] = id   //传进去以便分页跳转
	this.Data["pageIndex"] = pageIndex    //传进去用作是否active判断

	start := (pageIndex - 1)*pageSize


	//获取上一页页码
	prePage := pageIndex - 1
	if prePage <= 1{
		prePage = 1
	}
	this.Data["prePage"] = prePage

	//获取下一页页码
	nextPage := pageIndex + 1
	if nextPage > int(pageCount){
		nextPage = int(pageCount)
	}
	this.Data["nextPage"] = nextPage

	//按照一定顺序获取商品
	sort := this.GetString("sort")
	if sort == ""{
		o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",id).Limit(pageSize,start).All(&goods)  //显示部分商品（分页显示
		this.Data["sort"] = ""
		this.Data["goods"] = goods
	}else if sort == "price"{
		o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",id).OrderBy("Price").Limit(pageSize,start).All(&goods)
		this.Data["sort"] = "price"
		this.Data["goods"] = goods
	}else {
		o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",id).OrderBy("Sales").Limit(pageSize,start).All(&goods)
		this.Data["sort"] = "sale"
		this.Data["goods"] = goods
	}

	//返回视图
	this.TplName = "list.html"
}

//处理搜索
func(this*GoodsController)HandleSearch(){
	//获取数据
	goodsName := this.GetString("goodsName")
	o := orm.NewOrm()
	var goods []models.GoodsSKU
	//校验数据
		if goodsName == ""{
		o.QueryTable("GoodsSKU").All(&goods)
		this.Data["goods"] = goods
		ShowLaout(&this.Controller)
		this.TplName = "search.html"
		return
	}

	//处理数据  contains表包含某字段（具体参考讲义文档
	o.QueryTable("GoodsSKU").Filter("Name__icontains",goodsName).All(&goods)

	//返回视图
	this.Data["goods"] = goods
	ShowLaout(&this.Controller)  //布局
	this.TplName = "search.html"
}


