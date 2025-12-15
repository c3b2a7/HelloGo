package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	defer os.Remove("test.db")

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1) // 查询id为1的product
	fmt.Println(toJsonString(product))
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	fmt.Println(toJsonString(product))

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)
	fmt.Println(toJsonString(product))

	// 删除 - 删除product
	db.Delete(&product)
}

func toJsonString(v interface{}) string {
	js, _ := json.Marshal(v)
	return string(js)
}
