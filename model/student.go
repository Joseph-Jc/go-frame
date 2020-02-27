package model

type Student struct {
	Model
	Name      string `gorm:"type:varchar(5);not null;default:''" json:"name"`
	Age       int    `gorm:"type:smallint;not null;default:0" json:"age"`
	Classes   string `gorm:"type:varchar(255);not null;default:''" json:"classes"`
	Birthday  string `gorm:"type:date" json:"birthday"`
	SchoolDay string `gorm:"type:timestamp" json:"school_day"`
}

func CreateStudent() *Student {
	student := Student{Name: "张三", Age: 22, Classes: "计算机科学与技术1班", Birthday: "1997-01-01", SchoolDay: "2016-09-01 11:15:24"}
	DB.Create(&student)
	return &student
}
