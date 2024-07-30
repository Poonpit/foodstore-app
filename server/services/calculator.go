package services

import (
	// "fmt"
	// "errors"
	"server/repositories"
)

// CalculatorService interface for calculating total price
type CalculatorService interface {
	CalculateTotal(items map[string]int, hasMemberCard bool) (float64, error)
}

// calculatorService implements the CalculatorService interface
type calculatorService struct {
	menuRepo repositories.MenuRepository
}

// NewCalculatorService creates a new instance of CalculatorService
func NewCalculatorService(menuRepo repositories.MenuRepository) CalculatorService {
	return &calculatorService{menuRepo: menuRepo}
}

// CalculateTotal calculates the total price with discounts applied
func (s *calculatorService) CalculateTotal(items map[string]int, hasMemberCard bool) (float64, error) {
	menu := s.menuRepo.GetMenu()
	total := 0.0

	// Validate items and calculate total
	for itemName, quantity := range items {
		item, exists := menu[itemName]
		if !exists {
			return 0, ErrInvalidItem
		}
		total += item.Price * float64(quantity)

		// Apply bundle discount for specific sets
		if itemName == "Orange set" || itemName == "Pink set" || itemName == "Green set" {
			bundleDiscount := (quantity / 2) * int((item.Price*2)*0.05)
			total -= float64(bundleDiscount)
		}
	}
	// fmt.Println(total)

	// Apply member card discount (10%)
	if hasMemberCard {
		total *= 0.90
	}

	return total, nil
}
