package item

import "sync"

// TableRecord - Definisce il modulo per TableRecord
type TableRecord struct {
	isSelected bool
}

var tableRecord *TableRecord
var onceTableRecord sync.Once

// GetTableRecord - Restituisce l'unica istanza del modulo di table record
func GetTableRecord() *TableRecord {

	onceTableRecord.Do(func() {
		tableRecord = new(TableRecord)
	})

	return tableRecord
}

// Action - Definisce l'azione per il modulo
func (t TableRecord) Action() func() { return func() {} }

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
