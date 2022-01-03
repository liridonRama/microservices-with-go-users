package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/liridonrama/microservices-with-go-users/src/database"
	"github.com/liridonrama/microservices-with-go-users/src/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["passwordConfirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	isAmbassador, err := strconv.ParseBool(data["isAmbassador"])
	if err != nil {
		return fmt.Errorf("error while trying to parse isAmbassador to bool: %w", err)
	}

	fmt.Println(data)

	user := models.User{
		FirstName:    data["firstname"],
		LastName:     data["lastname"],
		Email:        data["email"],
		IsAmbassador: isAmbassador,
	}
	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}
