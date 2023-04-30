package model

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Chat    string `gorm:"not null"`
	Content string `gorm:"not null"`
}

// 创建记录
func CreateRecord(chat, content string) *gorm.DB {
	record := Record{Chat: chat, Content: content}
	return DB.Create(&record)
}

// 单聊查询历史聊天记录
func TwoFindAll(chat1 string, chat2 string, startTime time.Time, endTime time.Time) ([]Record, *gorm.DB) {
	var records []Record
	db := DB.Model(&Record{}).Where("(chat=? OR chat=?) AND created_at BETWEEN ? AND ?", chat1, chat2, startTime, endTime).Order("created_at DESC").Find(&records)
	return records, db
}
