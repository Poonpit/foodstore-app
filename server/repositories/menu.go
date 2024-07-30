package repositories

// Item struct representing a menu item
type Item struct {
	Name  string
	Price float64
}

// MenuRepository interface for accessing menu items
type MenuRepository interface {
	GetMenu() map[string]Item
}
