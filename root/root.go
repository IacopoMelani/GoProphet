package root

import (
	"fmt"

	"github.com/IacopoMelani/GoProphet/file"
)

var appName string

// GetAppName - Restituisce il nome dell'app
func GetAppName() string {

	if appName == "" {
		panic("Missing app name")
	}

	return appName
}

// InitRootDir si occupa di inizializzare la root principale
func InitRootDir() {

	fmt.Println("Choose the  app name: ")
	fmt.Scanf("%s", &appName)

	file.CheckAndCreateDir("./" + appName)
}
