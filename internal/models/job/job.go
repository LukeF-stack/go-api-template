package job

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Name    string
	Command string
	Args    string
}
