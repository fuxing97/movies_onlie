package PlayerController

import (
	mysql "movies_online/database"

	"movies_online/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SelByID(c *gin.Context) {
	var data []models.VideoList
	//获取路径参数
	ID := c.Query("id")
	//查询数据
	var total int
	
	mysql.DB().Where("movie_id = ?", ID).Order("id").Find(&data)
	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "没查询到数据",
		})
	} else {
		total = len(data)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": gin.H{
				"list":  data,
				"total": total,
			},
		})
	}
}

func SelSourceByID(c *gin.Context) {
	var data models.VideoInfo
	//获取路径参数
	ID := c.Query("id")
	mysql.DB().Where("id = ?", ID).Find(&data)
	if data.ID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "没查询到数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": data,
		})
	}
}
