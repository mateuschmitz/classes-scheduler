package classes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type ClassesHandler struct {
    repository *ClassesRepository
}

func (handler *ClassesHandler) GetAll(c *fiber.Ctx) error {
    var classes []Class = handler.repository.FindAll()
    return c.JSON(classes)
}

func (handler *ClassesHandler) Get(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    class, err := handler.repository.Find(id)

    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "status": 404,
            "message": "Class Not Found",
            "error":  err,
        })
    }

    return c.JSON(class)
}

func (handler *ClassesHandler) Create(c *fiber.Ctx) error {

    data := new(Class)
    
    if err := c.BodyParser(data); err != nil {
        return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Error parsing the payload", "error": err})
    }

    if (data.Name == "" || data.StartDate == "" || data.EndDate == "") {
        return c.Status(400).JSON(fiber.Map{
            "status":  400,
            "message": "Error parsing the payload",
            "error":   "Wrong parameters",
        })
    }

    class, err := handler.repository.Create(*data)

    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "status":  500,
            "message": "Error Creating Class",
            "error":   err,
        })
    }

    return c.JSON(class)
}

func NewClassesHandler(repository *ClassesRepository) *ClassesHandler {
    return &ClassesHandler{
        repository: repository,
    }
}

func Register(router fiber.Router, database *gorm.DB) {
    database.AutoMigrate(&Class{})
    classesRepository := NewClassesRepository(database)
    ClassesHandler := NewClassesHandler(classesRepository)
    
    classesRouter := router.Group("/classes")
    classesRouter.Get("/", ClassesHandler.GetAll)
    classesRouter.Get("/:id", ClassesHandler.Get)
    classesRouter.Post("/", ClassesHandler.Create)
}