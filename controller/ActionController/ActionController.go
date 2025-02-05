package ActionController

import (
	"log"
	mysql "movies_online/database"
	"strconv"

	"movies_online/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SelByVID(c *gin.Context) {
	var data models.VideoAction
	//获取路径参数
	ID, _ := strconv.Atoi(c.Query("vid"))
	//判断vid 是否存在
	mysql.DB().Select("id").Where("v_id = ?", ID).Find(&data)
	//判断id 是否存在
	if data.ID == 0 {
		data.PlayCount = 0
		data.VID = ID
		tx := mysql.DB().Create(&data)
		errCreate := tx.Error
		if errCreate != nil {
			log.Fatalf("create v_id action failed, err:%v", errCreate)
			return
		}
	}
	mysql.DB().Where("v_id = ?", ID).Find(&data)
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

func UpdatePlayCount(c *gin.Context) {
	var data models.VideoAction
	//获取路径参数
	ID, _ := strconv.Atoi(c.Query("vid"))
	//判断vid 是否存在
	mysql.DB().Select("id,play_count").Where("v_id = ?", ID).Find(&data)
	//判断id 是否存在
	if data.ID > 0 {
		playCount := data.PlayCount + 1
		mysql.DB().Where("id = ?", data.ID).Updates(
			models.VideoAction{
				PlayCount: playCount,
			},
		)
		c.JSON(http.StatusOK, gin.H{
			"code":      200,
			"msg":       "修改成功",
			"playcount": playCount,
		})
	}
	// mysql.DB().Where("v_id = ?", ID).Find(&data)
	// if data.ID <= 0 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 400,
	// 		"msg":  "没查询到数据",
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 200,
	// 		"msg":  "查询成功",
	// 		"data": data,
	// 	})
	// }
}

// func SelSourceByID(c *gin.Context) {
// 	var data models.VideoInfo
// 	//获取路径参数
// 	ID := c.Query("id")
// 	mysql.DB().Where("id = ?", ID).Find(&data)
// 	if data.ID <= 0 {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": 400,
// 			"msg":  "没查询到数据",
// 		})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": 200,
// 			"msg":  "查询成功",
// 			"data": data,
// 		})
// 	}
// }
