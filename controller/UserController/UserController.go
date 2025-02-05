package UserController

import (
	mysql "movies_online/database"

	"movies_online/models"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 增加用户
func Add(c *gin.Context) {
	//接值
	username := c.Request.FormValue("user_name")
	password := c.Request.FormValue("password")
	userType := c.Request.FormValue("user_type")
	gender := c.Request.FormValue("gender")
	email := c.Request.FormValue("email")
	//绑定模型
	user := models.User{
		UserName:     username,
		UserPassword: password,
		UserType:     userType,
		Gender:       gender,
		Email:        email,
	}
	mysql.DB().Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": "",
	})
}

/**
  删除
   1、找到对应id 的数据
   2、判断id是否存在
   3、成功 数据库删除  失败 返回id
*/

func Delete(c *gin.Context) {
	var data []models.User

	//接收id
	id := c.Param("id")

	//判断id 是否存在
	mysql.DB().Where("id = ?", id).Find(&data)

	//id 存在的情况则删除 不存在报错
	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "id没找到，删除失败",
		})
	} else {
		mysql.DB().Where("id = ?", id).Delete(&data)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	}

}

/**
修改
 1、找到对应id 所对应条目
 2、判断id是否存在
 3、如果存在修改条目 不存在返回id 没有找到
*/

func Update(c *gin.Context) {
	var data models.User

	//接收id
	id := c.Param("id")

	//判断id 是否存在
	mysql.DB().Select("id").Where("id = ?", id).Find(&data)

	//判断id 是否存在
	if data.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "id 没有找到",
		})
	} else {
		//接值
		username := c.Request.FormValue("user_name")
		password := c.Request.FormValue("password")
		email := c.Request.FormValue("email")

		mysql.DB().Where("id = ?", id).Updates(
			models.User{
				UserName:     username,
				UserPassword: password,
				Email:        email,
			},
		)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	}

}

/**
条件查询
*/

func Sel(c *gin.Context) {
	var data []models.User
	//获取路径参数
	name := c.Param("name")

	//查询数据
	mysql.DB().Where("user_name = ?", name).Find(&data)

	if len(data) == 0 {
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

func SelByID(c *gin.Context) {
	var data models.User
	//获取路径参数
	ID, _ := strconv.Atoi(c.Param("id"))
	//查询数据
	mysql.DB().Where("id = ?", ID).First(&data)
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

/**
分页查询
*/

func All(c *gin.Context) {
	var data []models.User
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	offsetVal := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offsetVal = -1
	}

	//总数
	var total int64
	//查询数据
	mysql.DB().Model(data).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&data)

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

// 用户登录
func UserLogin(c *gin.Context) {
	var data models.User
	//获取路径参数
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("passwd")
	if name == "" || password == "" {
		// 处理json 表单
		var userLogin models.LoginUser
		err := c.BindJSON(&userLogin)
		if err == nil {
			name = userLogin.Name
			password = userLogin.PassWord
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "json数据格式不正确",
			})
			return

		}
	}
	//查询数据
	mysql.DB().Where(&models.User{
		UserName:     name,
		UserPassword: password,
	}).Find(&data)
	if data.ID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "没查询到数据",
		})
	} else {
		data.UserPassword = ""
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "登入成功",
			"data": data,
		})
	}
}

// 增加普通用户
func UserAdd(c *gin.Context) {
	username := c.Request.FormValue("user_name")
	password := c.Request.FormValue("password")
	userType := c.Request.FormValue("user_type")
	gender := c.Request.FormValue("gender")
	email := c.Request.FormValue("email")
	isAdmin := "0"
	if username == "" {
		var userRegister models.RegisterUser
		err := c.BindJSON(&userRegister)
		if err == nil {
			username = userRegister.UserName
			password = userRegister.UserPassword
			userType = userRegister.UserType
			gender = userRegister.Gender
			email = userRegister.Email
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "json数据格式不正确",
			})
			return

		}
	}
	//绑定模型
	user := models.User{
		UserName:     username,
		UserPassword: password,
		UserType:     userType,
		Gender:       gender,
		Email:        email,
		IsAdmin:      isAdmin,
	}
	// 判断用户是否存在

	tx := mysql.DB().Create(&user)
	errCreate := tx.Error
	if errCreate != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   400,
			"msg":    "创建失败",
			"reason": errCreate,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": "",
	})
}
