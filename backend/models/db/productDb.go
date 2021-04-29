package db

import (
	"fmt"
	"backend/models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func open() *gorm.DB {
	USER := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "Shopping"
	CONNECT := USER + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&entity.Product{})

	fmt.Println("db connect: ", &db)
	return db
}

func FindAllProducts() []entity.Product {
	products := []entity.Product{}

	db := open()
	db.Order("ID asc").Find(&products)

	return products
}

func FindProduct(productId int) []entity.Product {
	product := []entity.Product{}

	db := open()
	db.First(&product, productId)

	return product
}

func InsertProduct(registerProduct *entity.Product) {
	db := open()

	db.Create(&registerProduct)
}

func UpdateStateProduct(productId int, productState int) {
	product := []entity.Product{}

	db := open()

	db.Model(&product).Where("ID = ?", productId).Update("State", productState)
}

func DeleteProduct(productId int) {
	product := []entity.Product{}

	db := open()

	db.Delete(&product, productId)
}
