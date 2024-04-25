package main

import (
	"github.com/Salonisaroha/user"
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to this journey")

}

func Routers(app *fiber.App) {
	app.Get("/user", user.GetUsers)
	app.Get("/users/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user", user.UpdateUser)

}

func main() {

	user.IntialMigration()
	app := fiber.New()
	app.Get("/", hello)
	Routers(app)
	app.Listen(":3000")

}
