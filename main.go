package main

import (
	"fmt"
	"log"

	"github.com/alaref-codes/subs/database"
	"github.com/alaref-codes/subs/subs"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Spec struct {
	Password string
}

func setupRoutes(app *fiber.App) {
	app.Get("/sub", subs.GetAllSubs)
	app.Get("/sub/:id", subs.GetOneSub)
	app.Post("/sub", subs.CreateSub)
	app.Delete("/sub/:id", subs.DeleteSub)
}

func initDatabase() {
	var err error
	pass := viper.Get("PASSWORD").(string)
	if err != nil {
		log.Fatal(err.Error())
	}
	dsn := "alaref:" + pass + "@tcp(127.0.0.1:3306)/subs?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection set")
}

func main() {
	viper.SetConfigFile(".env") // Reading the environment variable file
	viper.ReadInConfig()
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
