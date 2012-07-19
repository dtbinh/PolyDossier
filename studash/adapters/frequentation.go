package adapters

import (
	"studash/tools"
)

// ------------------------------------------------------ //
//            Builder pour la page frequentation          //
// ------------------------------------------------------ //
type FrequentationBuilder struct {
	DefaultAdapter
}

// Fonction GetParser de l'objet FrequantationBuilder
func (f FrequentationBuilder) GetParser() *tools.HTMLParser {
	f.parser = &tools.HTMLParser{"Frequentaion"}
	return f.parser
}
