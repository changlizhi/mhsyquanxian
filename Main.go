package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"mhsyquanxian/ziduan7skus"
	"mhsyquanxian/moxings"
	"time"
	"mhsyquanxian/gongjus"
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
			//获取到数据之后第一步，log下来gin自动做了
			//第二步验证授权
			//第三步验证字段是否合法
			//第四步拼接参数
			//第五步访问业务
			//第六步返回

			bbm := c.Param("Biaobianma")
			mx := moxings.Ziduan7s{
				Biaobianma:bbm,
			}
			ret := ziduan7skus.ChaxunZiduan7s(mx)
			c.JSON(http.StatusOK, ret)
			return
		})
		ziduan7s.POST("/liebiao/:Tiaojian", func(c *gin.Context) {
			tj := c.Param("Tiaojian")
			if tj == "quanbu"{
				ret := ziduan7skus.SuoyouZiduan7s()
				c.JSON(http.StatusOK, ret)
				return
			}
		})
	}
	r.Run(":8787")
}
