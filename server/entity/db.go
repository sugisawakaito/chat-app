package entity

import (
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

	PROTOCOL := "tcp(" + os.Getenv("PROTOCOL") + ")"
	dsn1 := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	Db, err = gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//PROTOCOL_REPLICA := "tcp(" + os.Getenv("PROTOCOL_REPLICA") + ")"
	dsnRep := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	err = Db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(dsnRep)},
		Policy:   dbresolver.RandomPolicy{},
	}))
	if err = autoMigrate(); err != nil {
		return
	}
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

func autoMigrate() (err error) {
	err = Db.AutoMigrate(&User{})
	return
}
