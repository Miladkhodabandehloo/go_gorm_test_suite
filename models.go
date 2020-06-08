package main

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	Title   string
	Content string
}

func GetModels() []interface{} {
	return []interface{}{&Blog{}}
}
