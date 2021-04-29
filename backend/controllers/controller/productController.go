package controller

import (
	"strconv"
	"backend/models/db"
	"backend/models/entity"

	"github.com/gin-gonic/gin"
)

const (
	NotPurchased = 0
	Purchased = 1
)

func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()

	c.JSON(200, resultProducts)
}

func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")

	productId, _ := strconv.Atoi(productIDStr)

	resultProduct := db.FindProduct(productId)

	c.JSON(200, resultProduct)
}

func AddProduct(c *gin.Context) {
	productName := c.PostForm("productName")
	productMemo := c.PostForm("productMemo")

	product := entity.Product{
		Name: productName,
		Memo: productMemo,
		State: NotPurchased,
	}

	db.InsertProduct(&product)
}

func ChangeStateProduct(c *gin.Context) {
	reqProductId := c.PostForm("productID")
	reqProductState := c.PostForm("productState")

	productId, _ := strconv.Atoi(reqProductId)
	productState, _ := strconv.Atoi(reqProductState)
	changeState := NotPurchased

	if productState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}

	db.UpdateStateProduct(productId, changeState)
}

func DeleteProduct(c *gin.Context) {
	productIDStr := c.PostForm("productID")
	productId, _ := strconv.Atoi(productIDStr)

	db.DeleteProduct(productId)
}
