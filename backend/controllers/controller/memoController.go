package controller

import (
	"backend/models/db"
	"backend/models/entity"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func FetchAllMemo(c *gin.Context) {
	resultMemos := db.FindAllMemos()

	c.JSON(200, resultMemos)
}

func FindMemo(c *gin.Context) {
	memoIDStr := c.Query("memoID")

	memoId, _ := strconv.Atoi(memoIDStr)

	resultMemo := db.FindMemo(memoId)

	c.JSON(200, resultMemo)
}

func AddMemo(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	memo := entity.Memo{
		Title: title,
		Content: content,
		Timestamp: time.Now(),
	}

	db.InsertMemo(&memo)
}

func UpdateMemo(c *gin.Context) {
	idStr := c.PostForm("id")
	content := c.PostForm("content")

	id, _ := strconv.Atoi(idStr)

	db.UpdateMemo(id, content)
}

func DeleteMemo(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)

	db.DeleteMemo(id)
}
