package entity

import "github.com/gofiber/fiber/v2"

type LogDetails struct {
	Message    string      `json:"message"`
	Reason     string      `json:"reason"`
	StatusCode int         `json:"status_code"`
	Request    interface{} `json:"request"`
	Response   interface{} `json:"response"`
	Context    *fiber.Ctx  `json:"-"`
}
