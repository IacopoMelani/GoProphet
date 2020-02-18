package item

import (
	"os"
	"sync"
)

// Exit - Definisce il modulo per uscire dal programma
type Exit struct{}

var exit *Exit
var onceExit sync.Once

// GetExit - Restituisce l'unica istanza del modulo di table record
func GetExit() *Exit {

	onceExit.Do(func() {
		exit = new(Exit)
	})

	return exit
}

// Action - Definisce l'azione per il modulo
func (e Exit) Action() func() {
	return func() {
		os.Exit(0)
	}
}

// IsSelected - Restituisce se il modulo è stato attualmente selezionato
func (e Exit) IsSelected() bool {
	return false
}

// Label - Restituisce la descrizione del modulo
func (e Exit) Label() string {
	return "Exit Prophet"
}

// Name - Restituisce il nome del modulo
func (e Exit) Name() string {
	return "exit-prophet"
}

// SetSelected - Imposta il modulo come selezionato
func (e Exit) SetSelected() {}
