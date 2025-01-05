package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"

	"authis/internal/config"
	"authis/internal/database"
	"authis/internal/models"
	"authis/internal/shared"
)

var (
	ctx = context.Background()
)

func Register(c *fiber.Ctx) error {
	data := map[string]string{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Username: data["username"],
		Password: shared.HashPassword(data["password"]),
	}

	database.GetDB().Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	data := map[string]string{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{}
	if err := user.FindByUsername(data["username"]); err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := shared.ComparePassword(user.Password, data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "wrong password",
		})
	}

	access, err := shared.AccessToken(user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	refresh, err := shared.RefreshToken(user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().GetRedisAddr(),
		Password: config.GetConfig().GetRedisPassword(),
		DB:       config.GetConfig().GetRedisDB(),
	})

	rdb.Set(strconv.Itoa(int(user.ID)), *access, time.Hour*24)
	rdb.Set("refresh_"+strconv.Itoa(int(user.ID)), *refresh, time.Hour*24*7)

	return c.JSON(fiber.Map{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func User(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	return c.JSON(user)
}
