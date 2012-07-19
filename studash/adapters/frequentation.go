package adapters

import (
	"studash/tools"
)

// ------------------------------------------------------ //
//            Builder pour la page frequentation          //
// ------------------------------------------------------ //
type FrequantationBuilder struct {
	DefaultAdapter
}

// Fonction GetParser de l'objet FrequantationBuilder
func (f FrequantationBuilder) GetParser() *tools.HTMLParser {
	f.parser = &tools.HTMLParser{"Frequentaion"}
	return f.parser
}
