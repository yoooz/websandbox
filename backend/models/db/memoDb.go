package db

import (
	"backend/models/entity"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openMemoDB() *gorm.DB {
	USER := "dev"
	PASSWORD := "dev"
	PROTOCOL := "tcp(127.0.0.1:4306)"
	DBNAME := "Dev"
	CONTENT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(CONTENT), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&entity.Memo{})

	fmt.Println("db connect: ", &db)

	return db
}

func FindAllMemos() []entity.Memo {
	memos := []entity.Memo{}

	db := openMemoDB()
	db.Order("ID asc").Find(&memos)

	return memos
}

func FindMemo(memoId int) []entity.Memo {
	memo := []entity.Memo{}

	db := openMemoDB()
	db.First(&memo, memoId)

	return memo
}

func InsertMemo(registerMemo *entity.Memo) {
	db := openMemoDB()

	db.Create(&registerMemo)
}

func UpdateMemo(memoId int, memoContent string) {
	memo := entity.Memo{}

	db := openMemoDB()

	db.Model(&memo).Where("ID = ?", memoId).Update("Content", memoContent)
}

func DeleteMemo(memoId int) {
	memo := []entity.Memo{}

	db := openMemoDB()

	db.Delete(&memo, memoId)
}
