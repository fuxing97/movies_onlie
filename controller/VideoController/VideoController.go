package VideoController

import (
	mysql "movies_online/database"
	"movies_online/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 分页查询
func All(c *gin.Context) {
	var data []models.MovieInfo
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	movieType := c.Query("movie_type")
	if movieType == "movie" {
		movieType = "电影"
	} else if movieType == "tv" {
		movieType = "电视剧"
	}

	if pageSize == 0 {
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	offsetVal := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offsetVal = -1
	}

	whereStr := "status = 2"
	if movieType != "" {
		whereStr += " and movie_type = '" + movieType + "'"
	}
	//总数
	var total int64
	//查询数据
	mysql.DB().Model(data).Where(whereStr).Count(&total).Order("update_time desc").Limit(pageSize).Offset(offsetVal).Find(&data)

	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "查询不到数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": gin.H{
				"list":     data,
				"total":    total,
				"pageNum":  pageNum,
				"pageSize": pageSize,
			},
		})
	}
}

func SelBymd5ID(c *gin.Context) {
	var data models.MovieInfo
	//获取路径参数
	movieID := c.Query("id")
	//查询数据
	mysql.DB().Where("movie_id = ? and status =2", movieID).First(&data)
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

// 模糊查找
func LikeSel(c *gin.Context){
	var data []models.MovieInfo
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	queryWord := c.Query("q")
	if pageSize == 0 {
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	offsetVal := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offsetVal = -1
	}

	whereStr := "status = 2"
	if queryWord != "" {
		whereStr += " and movie_name like '%" + queryWord + "%'"
	}
	//总数
	var total int64
	//查询数据
	mysql.DB().Model(data).Where(whereStr).Count(&total).Order("update_time desc").Limit(pageSize).Offset(offsetVal).Find(&data)

	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "查询不到数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": gin.H{
				"list":     data,
				"total":    total,
				"pageNum":  pageNum,
				"pageSize": pageSize,
			},
		})
	}
}

func SelByID(c *gin.Context) {
	var data models.MovieInfo
	//获取路径参数
	ID, _ := strconv.Atoi(c.Query("id"))
	//查询数据
	mysql.DB().Where("id = ? and status =2", ID).First(&data)
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

// 获取banner数据
func GetHot(c *gin.Context) {
	var data []models.MovieInfo
	mysql.DB().Model(data).Where("status =2 ").Order("update_time desc").Limit(5).Find(&data)
	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "查询不到数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": gin.H{
				"list": data,
			},
		})
	}
}


func Offline(c *gin.Context) {
	var data models.MovieInfo

	//接收id
	id := c.Param("movie_id")

	//判断id 是否存在
	mysql.DB().Select("id").Where("movie_id = ?", id).Find(&data)

	//判断id 是否存在
	if data.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "id 没有找到",
		})
	} else {
		mysql.DB().Where("movie_id = ?", id).Updates(
			models.MovieInfo{
				Status:     -1,
			},
		)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	}

}