package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NicName  string `db:"nic_name"`
	FullName string `db:"full_name"`
	Email    string `db:"email"`
}
