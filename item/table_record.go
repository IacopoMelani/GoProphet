package item

import (
	"os"
	"sync"

	"github.com/IacopoMelani/GoProphet/file"
)

// TableRecord - Definisce il modulo per TableRecord
type TableRecord struct {
	isSelected bool
}

const (
	modelsPath = "./models"
	tablePath  = modelsPath + "/table"

	urlTableRecordUsersFile  = "https://raw.githubusercontent.com/IacopoMelani/Go-Starter-Project/master/models/table/users.go"
	tableRecordUsersFilename = tablePath + "/users.go"

	urlTableRecordUsersTestFile  = "https://raw.githubusercontent.com/IacopoMelani/Go-Starter-Project/master/models/table/users_test.go"
	tableRecordUsersTestFilename = tablePath + "/users_test.go"
)

var (
	tableRecord     *TableRecord
	onceTableRecord sync.Once
)

func checkAndCreateDir(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

func initTableRecordDir() {

	checkAndCreateDir(modelsPath)
	checkAndCreateDir(tablePath)
	file.GetFileFromURL(urlTableRecordUsersFile, tableRecordUsersFilename)
	file.GetFileFromURL(urlTableRecordUsersTestFile, tableRecordUsersTestFilename)
}

// GetTableRecord - Restituisce l'unica istanza del modulo di table record
func GetTableRecord() *TableRecord {

	onceTableRecord.Do(func() {
		tableRecord = new(TableRecord)
	})

	return tableRecord
}

// Action - Definisce l'azione per il modulo
func (t TableRecord) Action() func() {

	return func() {
		initTableRecordDir()
	}
}

// IsSelected - Restituisce se il modulo è stato attualmente selezionato
func (t TableRecord) IsSelected() bool {
	return t.isSelected
}

// Label - Restituisce la descrizione del modulo
func (t TableRecord) Label() string {
	return "Table record"
}

// Name - Restituisce il nome del modulo
func (t TableRecord) Name() string {
	return "table-record"
}

// SetSelected - Imposta il modulo come selezionato
func (t *TableRecord) SetSelected() {
	t.isSelected = true
}
