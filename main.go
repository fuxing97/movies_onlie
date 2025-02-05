package main

import (
	"fmt"
	mysql "movies_online/database"
	"movies_online/routers"
)

func main() {
	
	// 初始化数据库
	fmt.Println("初始化数据库连接……")
	if mysql.SqlDb != nil {
		
		// 延迟退出关闭mysql
		defer mysql.SqlDb.Close()
	}
	router := routers.InitRouter()
	
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run(":8088")
}
