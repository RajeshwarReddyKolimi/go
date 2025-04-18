package main

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type Blog struct {
	gorm.Model
	Title    string `gorm:"not null, unique"`
	Content  string `gorm:"not null"`
	AuthorId uint
	Author   Author `gorm:"foreignKey:id;references:AuthorId"`
}
