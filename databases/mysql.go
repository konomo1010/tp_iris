package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"tp_iris/config"
)

var MySql *sql.DB

func init() {
	// 配置项
	fmt.Println(config.C.MySQL.User, config.C.MySQL.Passwd, &config.C.Log.Ac)

	username := config.C.MySQL.User
	password := config.C.MySQL.Passwd
	db := config.C.MySQL.Database
	host := config.C.MySQL.Host
	port := config.C.MySQL.Port

	dataSourceName := username + ":" + password + "@tcp" + "(" + host + ":" + port + ")/" + db + "?parseTime=true"

	var err error
	MySql, err = sql.Open(
		"mysql",
		dataSourceName,
	)
	if err != nil {
		log.Fatal("open : " + err.Error())
	}
	err = MySql.Ping()
	if err != nil {
		log.Fatal("ping : " + err.Error())
	}

	// mysql 的最大链接数设置
	MySql.SetMaxIdleConns(20)
	MySql.SetMaxOpenConns(20)

	fmt.Println("MySQL 初始化 完成!")
}
