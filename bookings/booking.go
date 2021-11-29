package bookings

import (
	"github.com/jinzhu/gorm"
)

type Booking struct {
    gorm.Model
    Name string `gorm:"type:varchar(100); Not Null" json:"name"`
    Date string `gorm:"type:date; Not Null" json:"date"`
    ClassID int `gorm:"type:int; Null"`
}