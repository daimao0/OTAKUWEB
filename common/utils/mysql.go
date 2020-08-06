package utils

import (
	_ "database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func InitDB() (db *sqlx.DB, err error) {
	username := beego.AppConfig.String("mysql_username")
	password := beego.AppConfig.String("mysql_password")
	host := beego.AppConfig.String("mysql_host")
	database := beego.AppConfig.String("mysql_database")
	dataSourceName := username + ":" + password + "@tcp(" + host + ")" + database
	db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatal("mysql连接出现错误，请检查dataSource是否正确", err.Error())
		return db, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return db, err
}
