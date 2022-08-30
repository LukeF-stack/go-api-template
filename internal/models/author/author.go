package author

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string `json:"name"`
}
