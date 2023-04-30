package service

import (
	"time"

	"github.com/chat/model"
)

func CreateRecord(chat, content string) error {
	err := model.CreateRecord(chat, content)
	if err.Error != nil {
		return err.Error
	} else {
		return nil
	}
}

func TwoFindAll(chat1, chat2 string, startTime, endTime time.Time) ([]model.Record, error) {
	records, err := model.TwoFindAll(chat1, chat2, startTime, endTime)
	if err.Error != nil {
		return nil, err.Error
	} else {
		return records, nil
	}
}
