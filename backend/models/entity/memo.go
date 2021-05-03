package entity

import "time"

type Memo struct {
	ID int `gorm:"primary_key;not null" json:"id"`
	Title string `gorm:"type:varchar(64);not null" json:"title"`
	Content string `gorm:"type:varchar(124)" json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

func (Memo) TableName() string {
	return "memo"
}
