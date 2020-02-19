package item

import (
	"github.com/IacopoMelani/GoProphet/menu"
	"github.com/IacopoMelani/GoProphet/root"
)

func getAppName() string {
	return "./" + root.GetAppName()
}

// ItitModules - Inizializza tutti i moduli
func ItitModules() {
	menu.RegisterItem(GetTableRecord())
	menu.RegisterItem(GetExit())
}
