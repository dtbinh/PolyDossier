package adapters

import (
	"studash/tools"
)

// ------------------------------------------------------ //
//               Builder pour la page finance             //
// ------------------------------------------------------ //
type FinanceBuilder struct {
	DefaultAdapter
}

// Fonction GetParser de l'objet FinanceBuilder
func (f FinanceBuilder) GetParser() *tools.HTMLParser {
	f.parser = &tools.HTMLParser{"Finance", nil, nil}
	return f.parser
}
