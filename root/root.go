package root

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/IacopoMelani/GoProphet/file"
)

var (
	appName  string
	rootPath string
	username string
)

// GetPathAppName - Restituisce il nome dell'app
func GetPathAppName() string {

	if username == "" {
		panic("missing username")
	}

	if appName == "" {
		panic("Missing app name")
	}

	rootPath = "github.com/" + username + "/" + appName

	return filepath.ToSlash(os.Getenv("GOPATH") + "/src/" + "github.com/" + username + "/" + appName)
}

// GetRootGoPath - Restituisce il percorso relativo del GOPATH per lo specifico progetto
func GetRootGoPath() string {
	return rootPath
}

// InitRootDir si occupa di inizializzare la root principale
func InitRootDir() {

	fmt.Print("Github username:")
	fmt.Scanf("%s\n", &username)

	fmt.Print("Choose the  app name: ")
	fmt.Scanf("%s\n", &appName)

	file.CheckAndCreateDir(GetPathAppName())
}
