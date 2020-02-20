package root

import (
	"fmt"
	"os"
	"strings"

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

	return strings.Replace(os.Getenv("GOPATH"), "\\", "/", 0) + "/src/" + "github.com/" + username + "/" + appName
}

// InitRootDir si occupa di inizializzare la root principale
func InitRootDir() {

	fmt.Print("Github username:")
	fmt.Scanf("%s\n", &username)

	fmt.Print("Choose the  app name: ")
	fmt.Scanf("%s\n", &appName)

	file.CheckAndCreateDir(GetPathAppName())
}
