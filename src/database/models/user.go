package models

type User struct {
	Id        string `gorm:"primaryKey"`
	Name      string `gorm:"size:30"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	DeletedAt int64  `gorm:"default:0"`
	Email     string
	Password  string
}
