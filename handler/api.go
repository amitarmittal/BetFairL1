package handler

import (
	"BetFairL1/dto"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Hello hanlde api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Simple Ping
// @Summary      Simple Ping
// @Description  Simple Ping to check the connection
// @Tags         Test
// @Accept       json
// @Produce      json
// @Success      200                  {object}  dto.CommonPortalRespDto
// @Failure      503                  {object}  dto.CommonPortalRespDto
// @Router       /test/simple-ping [post]
func SimplePing(c *fiber.Ctx) error {
	c.Accepts("json", "text")
	c.Accepts("application/json")
	t := time.Now().UnixNano()
	log.Println("SimplePing: nano time is - ", t)
	log.Println("SimplePing: milli time is - ", t/time.Hour.Milliseconds())
	respDto := new(dto.CommonPortalRespDto)
	respDto.Status = "RS_OK"
	respDto.ErrorDescription = ""
	return c.Status(fiber.StatusOK).JSON(respDto)
}
