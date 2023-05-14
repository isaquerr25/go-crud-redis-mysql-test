package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"

	"go-crud-redis-example/config"
	"go-crud-redis-example/database"
)

func main() {
	fmt.Println("Hello")
	
	//lets connecto the database mysql dan redis
	localConfig, err := config.LoadConfig(".")
	if err !=nil {
		log.Fatal("cannot load environtment variables", err)
		
	}
	
	
	// mysql
	db := database.ConnectionMySQLDb(&localConfig)
	rdb := database.ConnectionRedisDb(&localConfig)
	
	startServer(db,rdb)
	
	
	
}

func startServer( db *gorm.DB, rdb *redis.Client){


	app := fiber.New()
	
	err := app.Listen(":3400")
	
	if err != nil {
		panic(err)
	}
}