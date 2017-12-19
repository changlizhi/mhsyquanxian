package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mhsyquanxian/gongjus"
	"mhsyquanxian/moxings"
	"mhsyquanxian/zhanghao1syewus"
	"mhsyquanxian/ziduan7skus"
	"net/http"
	"os"
	"time"
)

type Retstruct struct {
	Mingcheng string
	Neirong   string
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin" + gongjus.Time2string(time.Now(), gongjus.NYRSFMXHX) + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.Use(Cors())
	ziduan7s := r.Group("ziduan7s")
	{
		ziduan7s.GET("/:Biaobianma", func(c *gin.Context) {
			//获取到数据之后第一步，log下来gin自动做了gin.log做了
			//第二步验证授权
			//第三步验证字段是否合法
			//第四步拼接参数
			//第五步访问业务
			//第六步返回

			bbm := c.Param("Biaobianma")
			mx := moxings.Ziduan7s{
				Biaobianma: bbm,
			}
			ret := ziduan7skus.ChaxunZiduan7s(mx)
			c.JSON(http.StatusOK, ret)
			return
		})
		ziduan7s.POST("/liebiao/:Tiaojian", func(c *gin.Context) {
			auth := c.GetHeader("Authorization")
			log.Println("auth-----", auth)
			claims := gongjus.Yanzhenglingpai(auth)
			log.Println("claims.Zhanghao----", claims.Zhanghao)
			tj := c.Param("Tiaojian")
			if tj == "quanbu" {
				ret := ziduan7skus.SuoyouZiduan7s()
				c.JSON(http.StatusOK, ret)
				return
			}
		})
	}
	zhanghao1s := r.Group("zhanghao1s")
	{
		zhanghao1s.POST("/denglu", func(c *gin.Context) {
			zh := c.PostForm("Zhanghao") //无论是用户名，手机号，邮箱都用一个输入框
			mm := c.PostForm("Mima")
			//通过业务层进行用户验证
			tg := zhanghao1syewus.YanzhengZhanghao1s(zh, mm)
			if tg {
				//如果验证通过
				lp, err := gongjus.Shengchenglingpai(zh)
				if err != nil {
					log.Println("生成令牌失败zhanghao1s/denglu---", err)
					c.JSON(http.StatusFailedDependency, gin.H{
						"err": "登录失败，请联系管理员",
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"Lingpai":  lp,
					"Ziyuan6s": "/ziduan7s/liebiao",
				})
				return
			}
		})
	}
	r.Run(":8787")
}
