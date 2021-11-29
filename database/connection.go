package database

import (
	"fmt"
	"strconv"

	"github.com/mateuschmitz/classes-scheduler/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ConnectDB() {
    var err error
    p := config.Config("DATABASE_PORT")
    port, err := strconv.ParseUint(p, 10, 32)

    configData := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.Config("DATABASE_HOST"),
        port,
        config.Config("DATABASE_USER"),
        config.Config("DATABASE_PASS"),
        config.Config("DATABASE_NAME"),
    )

    DB, err = gorm.Open(
        "postgres",
        configData,
    )

    if err != nil {
        fmt.Println(
            err.Error(),
        )
        panic("Failed to connect database.")
    }
}