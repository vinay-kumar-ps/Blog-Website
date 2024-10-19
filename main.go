package main

import (
	"log"
	"os"

	"github.com.vinay-kumar-ps/blogbackend/database"
	"github.com.vinay-kumar-ps/blogbackend/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
err :=godotenv.Load()
if err!=nil{
	log.Fatal("could not load port")
}
port :=os.Getenv("PORT")
app :=fiber.New()
route.SetUp(app)
app.Listen(":"+port)
}