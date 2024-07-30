package main

import (
	"fmt"
	"server/repositories"
	"server/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	// Mock menu repository
	menuRepo := repositories.NewMenuRepositoryMock()
	menu := map[string]repositories.Item{
		"Red set":    {Name: "Red set", Price: 50},
		"Green set":  {Name: "Green set", Price: 40},
		"Blue set":   {Name: "Blue set", Price: 30},
		"Yellow set": {Name: "Yellow set", Price: 50},
		"Pink set":   {Name: "Pink set", Price: 80},
		"Purple set": {Name: "Purple set", Price: 90},
		"Orange set": {Name: "Orange set", Price: 120},
	}
	menuRepo.On("GetMenu").Return(menu)

	calculatorService := services.NewCalculatorService(menuRepo)

	// POST /calculate
	app.Post("/calculate", func(c *fiber.Ctx) error {
		var req struct {
			Items         map[string]float64 `json:"items"`
			HasMemberCard bool               `json:"hasMemberCard"`
		}
		if err := c.BodyParser(&req); err != nil {
			fmt.Printf("Error parsing request body: %v", err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
		}

		items := make(map[string]int)
		for k, v := range req.Items {
			if int(v) > 0 {
				items[k] = int(v)
			}
		}

		total, err := calculatorService.CalculateTotal(items, req.HasMemberCard)
		if err != nil {
			fmt.Printf("Error calculating total: %v", err) 
			return c.Status(fiber.StatusInternalServerError).SendString("Error calculating total")
		}

		// fmt.Println("success")
		return c.JSON(fiber.Map{"total": total})
	})

	// Start the server
	app.Listen(":3000")
}
