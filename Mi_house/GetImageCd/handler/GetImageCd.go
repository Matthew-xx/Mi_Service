package handler

import (
	"Mi_house/ihomeWeb/utils"
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"

	"image/color"
	"time"

	GETIMAGECD "Mi_house/GetImageCd/proto/GetImageCd"
)

type GetImageCd struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetImageCd) Call(ctx context.Context, req *GETIMAGECD.Request, rsp *GETIMAGECD.Response) error {
	beego.Info("获取验证码图片 GetImageCode api/v1.0/imagecode/:uuid")
	// 生成验证码图片
	cap := captcha.New()

	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}
	//设置图片大小
	cap.SetSize(90, 41)
	//设置干扰强度
	cap.SetDisturbance(captcha.NORMAL)
	//设置前景色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	//设置背景色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	//生成随机验证码图片（4个数字
	CaptchaImage, str := cap.Create(4, captcha.NUM)
	beego.Info("验证码：", str)

	// 将UUID和验证码存入redis
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
	// 开始存入缓存：本次请求的uuid、本请求的答案，过期时间
	bm.Put(req.GetUuid(), str, time.Second*60)
	// 操作基本成功结束，开始返回正确结束信息
	beego.Info("图片验证码发送成功")
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)
	// 返回图片信息，注意图片是指针，接口传值不能传指针,需要解引用。
	image := *CaptchaImage
	img := *(image.RGBA)
	//返回图片拆分
	rsp.Pix = img.Pix
	rsp.Stride = int64(img.Stride)
	// rsp.Min.X = int64(img.Rect.Min.X)
	// rsp.Min.Y = int64(img.Rect.Min.Y)
	// rsp.Max.X = int64(img.Rect.Max.X)
	// rsp.Max.Y = int64(img.Rect.Max.Y)
	rsp.Min = &GETIMAGECD.Response_Point{X: int64(img.Rect.Min.X), Y: int64(img.Rect.Min.Y)}
	rsp.Max = &GETIMAGECD.Response_Point{X: int64(img.Rect.Max.X), Y: int64(img.Rect.Max.Y)}

	return nil
}

