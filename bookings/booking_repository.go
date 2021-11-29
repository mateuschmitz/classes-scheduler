package bookings

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type BookingsRepository struct {
    database *gorm.DB
}

func (repository *BookingsRepository) FindAll() []Booking {
    var bookings []Booking
    repository.database.Find(&bookings)
    return bookings
}

func (repository *BookingsRepository) Find(id int) (Booking, error) {
    var booking Booking
    err := repository.database.Find(&booking, id).Error
    if booking.Name == "" {
        err = errors.New("Booking not found")
    }
    return booking, err
}

func (repository *BookingsRepository) Create(booking Booking) (Booking, error) {
    err := repository.database.Create(&booking).Error
    if err != nil {
        return booking, err
    }

    return booking, nil
}

func NewBookingsRepository(database *gorm.DB) *BookingsRepository {
    return &BookingsRepository{
        database: database,
    }
}