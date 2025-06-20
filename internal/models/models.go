// name, surname, pantronymic, age, gender, nation(list)
package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"name"`
	Surname     string         `json:"surname" gorm:"surname"`
	Pantronymic string         `json:"pantronymic" gorm:"pantronymic"`
	Age         int            `json:"age" gorm:"age"`
	Gender      string         `json:"gender" gorm:"gender"`
	Nations     pq.StringArray `json:"nationality" gorm:"type:text[]"`
}

type RequestDTO struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Pantronymic string `json:"pantronymic"`
}
