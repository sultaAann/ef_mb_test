// name, surname, pantronymic, age, gender, nation(list)
package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Id          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name" gorm:"name"`
	Surname     string   `json:"surname" gorm:"surname"`
	Pantronymic string   `json:"pantronymic" gorm:"pantronymic"`
	Age         int      `json:"age" gorm:"age"`
	Gender      string   `json:"gender" gorm:"gender"`
	Nations     []string `json:"nationality" gorm:"nations"`
}
