package classes

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type ClassesRepository struct {
    database *gorm.DB
}

func (repository *ClassesRepository) FindAll() []Class {
    var classes []Class
    repository.database.Find(&classes)
    return classes
}

func (repository *ClassesRepository) Find(id int) (Class, error) {
    var class Class
    err := repository.database.Find(&class, id).Error
    if class.Name == "" {
        err = errors.New("Class not found")
    }
    return class, err
}

func (repository *ClassesRepository) FindByDate(date string) (Class, error) {
    var class Class
    err := repository.database.Where("start_date <= ? AND end_date >= ?", date, date).Find(&class).Error
    if class.Name == "" {
        err = errors.New("Class not found")
    }
    return class, err
}

func (repository *ClassesRepository) Create(class Class) (Class, error) {
    err := repository.database.Create(&class).Error
    if err != nil {
        return class, err
    }

    return class, nil
}

func NewClassesRepository(database *gorm.DB) *ClassesRepository {
    return &ClassesRepository{
        database: database,
    }
}