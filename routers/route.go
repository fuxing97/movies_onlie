package routers

import (
	"movies_online/controller/ActionController"
	"movies_online/controller/PlayerController"
	"movies_online/controller/UserController"
	"movies_online/controller/VideoController"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	route := gin.New()
	
	
	route.LoadHTMLGlob("templates/*")
	route.Static("/static", "static")
	// 创建日志文件
	logFile, err := os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// defer logFile.Close()
	// 创建一个logrus日志记录器
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})
	logger.SetOutput(logFile)
	// // 使用自定义的日志记录器替换默认的
	// route.Use(func(c *gin.Context) {
	// 	// 在请求开始时打印日志
	// 	logger.Infof("Request started: %s %s", c.Request.Method, c.Request.URL.Path)
	// 	// 继续执行后续的中间件或处理函数
	// 	c.Next()

	// 	// 在请求结束时打印日志
	// 	logger.Infof("Request finished: %s %s", c.Request.Method, c.Request.URL.Path)
	// })
	// route.Use(gin.Recovery()) // 使用 Recovery 中间件以捕获 panic
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	// 添加CORS中间件到Gin引擎
	// API路径
	APIgroup := route.Group("ajax")
	APIgroup.Use(cors.New(config))
	{
		userGroup := APIgroup.Group("user")
		{
			userGroup.POST("/add_user", UserController.UserAdd)
			userGroup.GET("/:id", UserController.SelByID)
			userGroup.POST("/login", UserController.UserLogin)

		}

		// 电影
		videoGroup := APIgroup.Group("movie")
		{
			videoGroup.GET("/getAll", VideoController.All)
			videoGroup.GET("/getMovieInfo", VideoController.SelByID)
			videoGroup.GET("/getMd5IDInfo", VideoController.SelBymd5ID)
			videoGroup.GET("/getbanner", VideoController.GetHot)
			videoGroup.GET("/offline/:movie_id", VideoController.Offline)
			videoGroup.GET("/query", VideoController.LikeSel)
		}

		// 播放页数据
		playerGroup := APIgroup.Group("player")
		{
			playerGroup.GET("/getVideoInfo", PlayerController.SelByID)
			playerGroup.GET("/getSource", PlayerController.SelSourceByID)

		}

		videoActionGroup := APIgroup.Group("video_action")
		{
			videoActionGroup.GET("/get", ActionController.SelByVID)
			videoActionGroup.GET("/add_play", ActionController.UpdatePlayCount)

		}
	}

	// 首页
	route.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	// 详情页
	detailPageGroup := route.Group("voddetail")
	{
		detailPageGroup.GET("/:id.html", func(c *gin.Context) {
			c.HTML(http.StatusOK, "detail.html", nil)
		})

	}

	// 播放页
	route.GET("/player/:id.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "player.html", nil)
	})

	route.GET("/movie_list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", nil)
	})

	route.GET("/tv_list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", nil)
	})

	route.GET("/tv_list/page/:pageNum", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", nil)
	})
	route.GET("/movie_list/page/:pageNum", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", nil)
	})

	route.GET("/login/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})

	route.GET("/source/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "source_list.html", nil)
	})

	route.GET("/source/list/page/:pageNum", func(c *gin.Context) {
		c.HTML(http.StatusOK, "source_list.html", nil)
	})
	route.GET("/query/:queryword", func(c *gin.Context) {
		c.HTML(http.StatusOK, "query.html", nil)
	})
	route.GET("/query/:queryword/page/:pageNum", func(c *gin.Context) {
		c.HTML(http.StatusOK, "query.html", nil)
	})
	return route
}
