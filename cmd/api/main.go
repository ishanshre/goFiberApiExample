package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ishanshre/goFiberApiExample/internals/config"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/goFiberApiExample/internals/driver"
	"github.com/ishanshre/goFiberApiExample/internals/handler"
	"github.com/joho/godotenv"
)

var port string = ":8000"
var global config.AppConfig // gloabel config
var connString string
var database string = "postgres"

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error in loading env files: %s", err.Error())
	}

	db, err := Run()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.SQL.Close()

	app := fiber.New()

	Router(&global, app)

	app.Listen(port)
}

func Run() (*driver.DB, error) {
	global.InProduction = false

	log.Println("connecting to database")
	connString = os.Getenv(database)
	db, err := driver.ConnectSQL(database, connString)
	if err != nil {
		return nil, fmt.Errorf("error in connecting to dataase :%s", err.Error())
	}
	log.Println("connected to database")

	repo := handler.NewRepo(&global, db)
	handler.NewHandler(repo)
	return db, nil
}
