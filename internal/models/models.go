// name, surname, pantronymic, age, gender, nation(list)
package models

type Person struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Surname     string   `json:"surname"`
	Pantronymic string   `json:"pantronymic"`
	Age         int      `json:"age"`
	Gender      string   `json:"gender"`
	Nations     []string `json:"nationality"`
}
