package menu

import (
	"errors"
	"sync"

	"github.com/inancgumus/screen"
	"github.com/manifoldco/promptui"
)

// Item -
type Item interface {
	Name() string
	Label() string
	Action() func()
	IsSelected() bool
	SetSelected()
}

var (
	mu   = &sync.Mutex{}
	menu []Item
)

func cleanScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

// getItemByName - Restituisce l'item dato il nome
func getItemByName(name string) (Item, error) {

	for _, i := range menu {
		if i.Name() == name {
			return i, nil
		}
	}
	return nil, errors.New("Item non trovato")
}

// RegisterItem - Si occcupa di registrare un nuovo elemento del menù
func RegisterItem(i Item) {
	mu.Lock()
	menu = append(menu, i)
	mu.Unlock()
}

// GetItemForSelect  - Restituisce i nome degli item per la select
func GetItemForSelect() []string {

	var items []string

	mu.Lock()
	for _, i := range menu {
		if !i.IsSelected() {
			items = append(items, i.Name())
		}
	}
	mu.Unlock()

	return items
}

// ShowMenu - Si occupa di mostrare il menù di selezione
func ShowMenu() {

	for {

		cleanScreen()

		prompt := promptui.Select{
			Label: "Select module",
			Items: GetItemForSelect(),
		}

		_, result, err := prompt.Run()

		if err != nil {
			panic(err)
		}

		item, err := getItemByName(result)
		if err != nil {
			panic(err)
		}

		item.SetSelected()
		fn := item.Action()
		fn()
	}
}
