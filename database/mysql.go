package mysql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var SqlDb *sql.DB

func DB() *gorm.DB {
	// 从配置文件加载配置
	data, err := ioutil.ReadFile("conf/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse YAML file: %v", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)
	sqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("数据库mysql连接异常:错误详情:%v", err)
	}
	sqlDB, errDB := sqlDb.DB()
	if err != nil {
		log.Fatalf("生成sqlDB实例错误, 详情:%v", errDB)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(1 * time.Minute)
	return sqlDb

}
