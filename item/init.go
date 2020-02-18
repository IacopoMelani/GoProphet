package item

import (
	"github.com/IacopoMelani/GoProphet/menu"
)

// ItitModules - Inizializza tutti i moduli
func ItitModules() {
	menu.RegisterItem(GetTableRecord())
	menu.RegisterItem(GetExit())
}
