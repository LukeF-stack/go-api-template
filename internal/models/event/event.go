package event

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name    string
	Message string
}
