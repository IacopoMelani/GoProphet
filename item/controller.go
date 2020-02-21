package item

import (
	"io/ioutil"
	"strings"
	"sync"

	"github.com/IacopoMelani/GoProphet/root"

	"github.com/IacopoMelani/GoProphet/file"
)

// Controller - Modulo per la gestione di services
type Controller struct {
	isSelected bool
}

const (
	controllersPath = "/controllers"
	routesPath      = "/routes"

	urlControllerFile  = "https://raw.githubusercontent.com/IacopoMelani/Go-Starter-Project/master/controllers/controller.go"
	controllerFilename = controllersPath + "/controller.go"

	urlControllerTestFile  = "https://raw.githubusercontent.com/IacopoMelani/Go-Starter-Project/master/controllers/controller_test.go"
	controllerTestFilename = controllersPath + "/controller_test.go"

	urlRoutesFile  = "https://raw.githubusercontent.com/IacopoMelani/Go-Starter-Project/master/routes/routes.go"
	routesFilename = routesPath + "/routes.go"

	urlRoutesTestFile  = "https://raw.githubusercontent.com/IacopoMelani/Go-Starter-Project/master/controllers/routes_test.go"
	routesTestFilename = routesPath + "/routes_test.go"

	routesInternalDependeciesControllers = "github.com/IacopoMelani/Go-Starter-Project/controllers"
)

var (
	controller     *Controller
	onceController sync.Once
)

func initControllersDir() {

	file.CheckAndCreateDir(GetPathAppName() + controllersPath)
	file.CheckAndCreateDir(GetPathAppName() + routesPath)

	file.GetFileFromURL(urlControllerFile, GetPathAppName()+controllerFilename)
	file.GetFileFromURL(urlControllerTestFile, GetPathAppName()+controllerTestFilename)
	file.GetFileFromURL(urlRoutesFile, GetPathAppName()+routesFilename)
	file.GetFileFromURL(urlRoutesTestFile, GetPathAppName()+routesTestFilename)

	read, err := ioutil.ReadFile(GetPathAppName() + routesFilename)
	if err != nil {
		panic(err)
	}

	newContent := strings.Replace(string(read), routesInternalDependeciesControllers, root.GetRootGoPath()+"/controllers", -1)
	err = ioutil.WriteFile(GetPathAppName()+routesFilename, []byte(newContent), 0)
	if err != nil {
		panic(err)
	}
}

// GetController - Restituisce l'istanza di controller
func GetController() *Controller {

	onceController.Do(func() {
		controller = new(Controller)
	})

	return controller
}

// Action - Denifisce l'azione per il modulo
func (c Controller) Action() func() {

	return func() {
		initControllersDir()
	}
}

// IsSelected - Restituisce se il modulo è stato attualmente selezionato
func (c Controller) IsSelected() bool {
	return c.isSelected
}

// Label - Restituisce la descrizione del modulo
func (c Controller) Label() string {
	return "Route & Controller"
}

// Name - Restituisce il nome del modulo
func (c Controller) Name() string {
	return "route-controller"
}

// SetSelected - Imposta il modulo come selezionato
func (c *Controller) SetSelected() {
	c.isSelected = true
}
