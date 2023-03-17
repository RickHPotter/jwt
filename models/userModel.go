package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // [1]
	Email      string `gorm:"unique"`
	Password   string
}

// ! [1]
// ! gorm.Model is a call to the following struct
// type Model struct {
// 	ID        uint `gorm:"primarykey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt DeletedAt `gorm:"index"`
// }
//
// ! In the same manner, we could also use mapping attributes to only get what we need
// type User2 struct {
// 	ID uint `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	Email    string `gorm:"unique"`
// 	Password string
// }
