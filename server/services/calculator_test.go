package services_test

import (
	// "errors"
	// "fmt"
	"server/repositories"
	"server/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotal(t *testing.T) {
	type testCase struct {
		name          string
		items         map[string]int
		hasMemberCard bool
		expected      float64
		err           error
	}

	cases := []testCase{
		{name: "basic", items: map[string]int{"Red set": 1, "Green set": 1}, hasMemberCard: false, expected: 90},                         // no discount case
		{name: "member card discount", items: map[string]int{"Red set": 1, "Green set": 1}, hasMemberCard: true, expected: 81},           // 10% discount case
		{name: "bundle discount", items: map[string]int{"Orange set": 5}, hasMemberCard: false, expected: 576},                           // 5% off for 2 pairs case
		{name: "bundle discount with member card", items: map[string]int{"Orange set": 5}, hasMemberCard: true, expected: 518.4},         // 5% off for 2 pairs + 10% member card discount case
		{name: "invalid item", items: map[string]int{"Invalid set": 1}, hasMemberCard: false, expected: 0, err: services.ErrInvalidItem}, // not in menu case
	}

	menu := map[string]repositories.Item{
		"Red set":    {Name: "Red set", Price: 50},
		"Green set":  {Name: "Green set", Price: 40},
		"Blue set":   {Name: "Blue set", Price: 30},
		"Yellow set": {Name: "Yellow set", Price: 50},
		"Pink set":   {Name: "Pink set", Price: 80},
		"Purple set": {Name: "Purple set", Price: 90},
		"Orange set": {Name: "Orange set", Price: 120},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Arrange
			menuRepo := repositories.NewMenuRepositoryMock()
			menuRepo.On("GetMenu").Return(menu)

			calculatorService := services.NewCalculatorService(menuRepo)

			// Act
			total, err := calculatorService.CalculateTotal(c.items, c.hasMemberCard)

			// Assert
			if c.err != nil {
				assert.ErrorIs(t, err, c.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.expected, total)
			}
		})
	}
}
