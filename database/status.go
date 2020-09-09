package database

import (
	"log"
	"time"
)

type (
	status struct {
		Id     int
		Status int
		Date   time.Time
	}

	count struct {
		Status int
		Count  int
	}
)

func NewStatus(code int) *status {
	s := &status{
		Status: code,
		Date:   time.Now().UTC(),
	}
	return s
}

func (s *status) Set() error {
	err := db.Model(&status{}).Save(s).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (Database) GetCount(date time.Time) ([]*count, error) {
	result := []*count{}
	err := db.Model(&status{}).
		Select("status, count(status) as count").
		Where("date > ?", date).
		Group("status").
		Find(&result).Error
	return result, err
}
