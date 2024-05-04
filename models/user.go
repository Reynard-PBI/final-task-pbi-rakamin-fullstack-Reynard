package models

import (
	// "time"
	// "gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID         int     `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Email      string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password   string    `gorm:"type:varchar(255);not null" json:"password"`
	// Photo      Photo     `gorm:"foreignKey:UserID" json:"photo"`
	Photos 	   []Photo 	 `gorm:"foreignKey:UserID"`
	Created_At string `gorm:"type:timestamp;default:CURRENT_TIMESTAMP()" json:"created_at"`
	Updated_At string `gorm:"type:timestamp;default:CURRENT_TIMESTAMP()" json:"updated_at"`
}