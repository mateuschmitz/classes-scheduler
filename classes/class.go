package classes

import (
	"github.com/jinzhu/gorm"
)

type Class struct {
    gorm.Model
    Name      string `gorm:"type:varchar(100); Not Null" json:"name"`
    StartDate string `gorm:"type:date; Not Null" json:"start_date"`
    EndDate   string `gorm:"type:date; Not Null" json:"end_date"`
	Capacity  int `gorm:"Not Null" json:"capacity"`
}