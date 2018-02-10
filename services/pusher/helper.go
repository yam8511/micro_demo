package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Usage 顯示說明
func Usage() {
	if GetAppEnv() == "" {
		fmt.Printf(`
			📖  Micro 的 Pusher 說明 📖
			請傳入以下環境變數：

			⚙  PROJECT_ENV : 專案環境
				✏  dev 開發
				✏  prod 正式

			📌 舉例：  PROJECT_ENV=dev ./pusher
`)
		os.Exit(0)
	} else {
		log.Printf("⚙  PROJECT_ENV: %s", GetAppEnv())
	}
}

// GetAppRoot 取執行檔的根目錄
func GetAppRoot() string {
	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println("🎃 GetAppRoot：取根目錄失敗，自動抓取 PROJECT_ROOT 的環境變數🎃")
		return os.Getenv("PROJECT_ROOT")
	}
	return root
}

// GetAppEnv 取環境變數
func GetAppEnv() string {
	return os.Getenv("PROJECT_ENV")
}

// CheckDangerError 確認危險錯誤，若錯誤直接死亡
func CheckDangerError(err error) {
	if err != nil {
		log.Println(err)
		log.Fatalln("👽  Pusher Service Dead 👽")
	}
}

// UniqueStringSlice 將陣列Unique
func UniqueStringSlice(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
