package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func apiRoutesV1(rootGroup fiber.Router) {
	v1Group := rootGroup.Group("/v1")
	fmt.Print(v1Group) //nolint:forbidigo
}
