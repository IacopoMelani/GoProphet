package menu

import (
	"sync"
)

// Item -
type Item interface {
	Name() string
	Label() string
	Action() func()
}

var (
	mu   sync.Mutex
	menu []Item
)

// RegisterItem - Si occcupa di registrare un nuovo elemento del menù
func RegisterItem(i Item) {
	mu.Lock()
	menu = append(menu, i)
	mu.Unlock()
}

// GetItemForSelect  - Restituisce i nome degli item per la select
func GetItemForSelect() []string {

	var items []string

	for _, i := range menu {
		items = append(items, i.Name())
	}

	return items
}
