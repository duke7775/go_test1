package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	ID   int
	Name string
	Age  int
}

var db *gorm.DB

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	fmt.Println("MYSQL successful")

	db.AutoMigrate(&Person{})

	db.Create(&Person{ID: 681200, Name: "jack", Age: 20})
	db.Create(&Person{ID: 681201, Name: "tom", Age: 28})

	var people []Person
	db.Find(&people)
	for _, p := range people {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", p.ID, p.Name, p.Age)
	}
}
