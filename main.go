package main

import (
	"github.com/IacopoMelani/GoProphet/item"
	"github.com/IacopoMelani/GoProphet/menu"
	"github.com/IacopoMelani/GoProphet/root"
)

func main() {

	item.ItitModules()

	root.InitRootDir()

	menu.ShowMenu()
}
