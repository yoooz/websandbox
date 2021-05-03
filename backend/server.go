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

	memo := router.Group("memo")
	{
		memo.GET("/fetchAll", controller.FetchAllMemo)
		memo.GET("/find", controller.FindMemo)
		memo.POST("/add", controller.AddMemo)
		memo.POST("/update", controller.UpdateMemo)
		memo.POST("/delete", controller.DeleteMemo)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
