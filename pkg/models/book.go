package models

// Book , Creating model that will handel database
type Book struct {
	Id     int    `json:"id" gorm:"primaryKey""`
	Title  string `json:"title"`
	Author string ` json:"author"`
	Desc   string `json:"desc"`
}
