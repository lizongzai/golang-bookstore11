package model

type Book struct {
	Id     int64  `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"not null" binding:"required"`
	Author string `json:"author" gorm:"not null" binding:"required"`
	Users  []User `gorm:"many2many:book_users"`
}

func (Book) TableName() string {
	return "book"
}
