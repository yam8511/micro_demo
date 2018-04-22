package main

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SetUpTable 設置資料表
func SetUpTable() (err error) {
	err = OrmWriteHandle(func(db *gorm.DB) (handleErr error) {
		db.AutoMigrate(&User{})
		return
	})
	return
}

func getConnectionInfo() string {
	user := "root"
	password := "qwe123"
	dbName := "Demo"
	dbHost := "0.0.0.0"
	dbPort := ":3306"
	loc := url.QueryEscape("Asia/Taipei")
	connectionName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=%s", user, password, dbHost+dbPort, dbName, loc)
	return connectionName
}

// OrmWriteHandle 資料庫進行寫的處理
func OrmWriteHandle(handle func(db *gorm.DB) (handleErr error)) (ormErr error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			ormErr = fmt.Errorf("%v", catchErr)
		}
	}()

	// db, err := gorm.Open("sqlite3", "./auth.db")
	db, err := gorm.Open("mysql", getConnectionInfo())
	if err != nil {
		ormErr = err
		return
	}
	defer db.Close()
	db.LogMode(true)
	ormErr = handle(db)

	return
}

// OrmReadHandle 資料庫進行讀的處理
func OrmReadHandle(handle func(db *gorm.DB) (handleErr error)) (ormErr error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			ormErr = fmt.Errorf("%v", catchErr)
		}
	}()

	// db, err := gorm.Open("sqlite3", "./auth.db")
	db, err := gorm.Open("mysql", getConnectionInfo())
	if err != nil {
		ormErr = err
		return
	}
	defer db.Close()
	db.LogMode(true)
	ormErr = handle(db)

	return
}
