package main

import (
	"encoding/json"
	"fmt"
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
	totalWorkouts, err = du.ReadTotalWorkoutFile()
	if err != nil {
		log.Fatal(err)
	}

	server.Get("/all", func(c *fiber.Ctx) error {
		return c.JSON(totalWorkouts)
	})

	server.Post("/all/mutate", func(c *fiber.Ctx) error {
		var (
			reqBody []byte
			errMsg  string
		)

		reqBody = c.Body()

		err = json.Unmarshal(reqBody, &totalWorkouts)
		if err != nil {
			errMsg = fmt.Sprintf("%v", err)
			return c.SendString(errMsg)
		}

		err = du.WriteTotalWorkout(reqBody)
		if err != nil {
			errMsg = fmt.Sprintf("%v", err)
			return c.SendString(errMsg)
		}

		return c.SendString("success")
	})

	server.Get("/:weekNum/:participant?/:exercise?", func(c *fiber.Ctx) error {
		var (
			errMsg       string
			weekNum, err = c.ParamsInt("weekNum")
			participant  = c.Params("participant")
			exercise     = c.Params("exercise")
		)
		if err != nil {
			errMsg = fmt.Sprintf("%v", err)
			return c.SendString(errMsg)
		}

		if weekNum < 0 || weekNum == len(totalWorkouts) {
			return c.SendString("weekNum parameter is out of range")
		}

		if participant == "" {
			targetData := totalWorkouts[weekNum]

			return c.JSON(targetData)
		} else {
			targetData := totalWorkouts[weekNum].Participants[participant]
			if exercise == "" {
				return c.JSON(targetData)
			} else {
				return c.JSON(targetData[exercise])
			}
		}
	})

	server.Get("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(server.Listen(":3000"))
}
