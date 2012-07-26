package adapters

import (
	"studash/tools"
)

// ------------------------------------------------------ //
//           Builder pour la page Renseignements          //
// ------------------------------------------------------ //
type InformationBuilder struct {
	DefaultAdapter
}

// Fonction GetParser de l'objet InformationBuilder
func (i InformationBuilder) GetParser() *tools.HTMLParser {
	i.parser = &tools.HTMLParser{"Information", nil}
	return i.parser
}
