package bookings

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/mateuschmitz/classes-scheduler/classes"
)

type BookingsHandler struct {
    repository *BookingsRepository
    classesRepository *classes.ClassesRepository
}

func (handler *BookingsHandler) GetAll(c *fiber.Ctx) error {
    var bookings []Booking = handler.repository.FindAll()
    return c.JSON(bookings)
}

func (handler *BookingsHandler) Get(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    booking, err := handler.repository.Find(id)

    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "status": 404,
            "error":  err,
        })
    }

    return c.JSON(booking)
}

func (handler *BookingsHandler) Create(c *fiber.Ctx) error {

    data := new(Booking)
    
    if err := c.BodyParser(data); err != nil {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error parsing the payload", "error": err})
    }

    if (data.Name == "" || data.Date == "") {
        return c.Status(400).JSON(fiber.Map{
            "status":  400,
            "message": "Error parsing the payload",
            "error":   "Wrong parameters",
        })
    }

    class, err := handler.classesRepository.FindByDate(data.Date)
    if err != nil {
        return c.Status(406).JSON(fiber.Map{
            "status":  406,
            "message": "Class not found",
            "error":   "There's no classes for this day",
        })
    }

    data.ClassID = int(class.ID)
    item, err := handler.repository.Create(*data)

    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "status":  500,
            "message": "Failed creating item",
            "error":   err,
        })
    }

    return c.JSON(item)
}

func NewBookingsHandler(repository *BookingsRepository, classesRepository *classes.ClassesRepository) *BookingsHandler {
    return &BookingsHandler{
        repository: repository,
        classesRepository: classesRepository,
    }
}

func Register(router fiber.Router, database *gorm.DB) {
    database.AutoMigrate(&Booking{})
    BookingsRepository := NewBookingsRepository(database)
    ClassesRepository := classes.NewClassesRepository(database)
    BookingsHandler := NewBookingsHandler(BookingsRepository, ClassesRepository)

    bookingsRouter := router.Group("/bookings")
    bookingsRouter.Get("/", BookingsHandler.GetAll)
    bookingsRouter.Get("/:id", BookingsHandler.Get)
    bookingsRouter.Post("/", BookingsHandler.Create)
}