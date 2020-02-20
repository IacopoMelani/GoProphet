package item

import (
	"github.com/IacopoMelani/GoProphet/menu"
	"github.com/IacopoMelani/GoProphet/root"
)

// GetPathAppName - Restituisce il path completo della root del progetto
func GetPathAppName() string {
	return root.GetPathAppName()
}

// ItitModules - Inizializza tutti i moduli
func ItitModules() {
	menu.RegisterItem(GetTableRecord())
	menu.RegisterItem(GetExit())
}
