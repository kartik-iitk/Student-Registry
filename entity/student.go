package entity

//Student object for REST(CRUD)
type Student struct {
	ID     int    `json:"ID" gorm:"primaryKey"`
	RollNo int    `json:"rollNo" gorm:"unique"`
	Name   string `json:"name"`
	Branch string `json:"branch"`
	UserId int    `json:"userId"`
}
