package main

import (
	"log"
	"backend/controllers/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	serve()
}

func serve() {
	router := gin.Default()

	router.GET("/fetchAllProducts", controller.FetchAllProducts)
	router.GET("/fetchProduct", controller.FindProduct)
	router.POST("/addProduct", controller.AddProduct)
	router.POST("/changeStateProduct", controller.ChangeStateProduct)
	router.POST("/deleteProduct", controller.DeleteProduct)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
