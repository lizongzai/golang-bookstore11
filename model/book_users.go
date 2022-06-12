package model

type book_users struct {
	UserId int64 `json:"userid" gorm:"primaryKey"`
	BookId int64 `json:"bookid" gorm:"primaryKey"`
}
