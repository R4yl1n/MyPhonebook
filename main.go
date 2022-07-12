package main

import (
	"log"

	"github.com/R4yl1n/MariaDB/models"
	"github.com/gofiber/fiber/v2"
)

var responseDBA = make([]models.MariaDB, 0) //response as an Array
var resp models.MariaDB                     //response as single struct

func main() {

	app := fiber.New()

	app.Get("api/datas", func(c *fiber.Ctx) error {
		responseDBA = models.Getall()

		return c.JSON(responseDBA)
	})

	app.Post("api/datas", func(c *fiber.Ctx) error {
		if err := c.BodyParser(&resp); err != nil {
			return err
		}
		log.Print(resp)

		return c.JSON(models.Inserdata(resp))

	})

	app.Patch("api/datas", func(c *fiber.Ctx) error {

		if err := c.BodyParser(&resp); err != nil {
			return err
		}
		log.Print(resp)

		return c.JSON((models.Replacenumber(resp)))
	})

	app.Delete("api/datas", func(c *fiber.Ctx) error {
		if err := c.BodyParser(&resp); err != nil {
			return err
		}

		return c.SendString(models.Deletedatas(resp))
	})

	log.Fatal(app.Listen(":4000"))
}
