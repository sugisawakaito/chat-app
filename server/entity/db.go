package entity

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var (
	Db  *gorm.DB
	err error
)

func DBConnect() {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASS")
	DBNAME := os.Getenv("MYSQL_DATABASE")
	PROTOCOL := fmt.Sprintf("tcp(%s)", os.Getenv("PROTOCOL"))

	dsn1 := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	Db, err = gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	dsnRep := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	err = Db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(dsnRep)},
		Policy:   dbresolver.RandomPolicy{},
	}))
}

func GetDB() *gorm.DB {
	return Db
}

func Close() {
	db, _ := Db.DB()
	if err = db.Close(); err != nil {
		return
	}
}
