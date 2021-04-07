package main

import (
	"encoding/json"
	"errors"
	"log"

	du "github.com/branogarbo/rhsJudoAPI/dataUtil"
	"github.com/gofiber/fiber/v2"
)

var (
	server        *fiber.App
	totalWorkouts du.TotalWorkouts
	err           error
)

func main() {
	server = fiber.New()
	totalWorkouts, err = du.ReadTotalWorkoutStruct()
	if err != nil {
		log.Fatal(err)
	}

	server.Get("/all", func(c *fiber.Ctx) error {
		return c.JSON(totalWorkouts)
	})

	server.Post("/all", func(c *fiber.Ctx) error {
		var (
			reqBody []byte
		)

		reqBody = c.Body()

		if !json.Valid(reqBody) {

			return errors.New("incoming data is invalid")
		}

		err = du.WriteTotalWorkoutStruct(reqBody)
		if err != nil {
			return err
		}

		return c.SendString("success")
	})

	log.Fatal(server.Listen(":3000"))
}
