package adapters

import (
	"studash/tools"
)

// ------------------------------------------------------ //
//              Builder pour la page Bulletin             //
// ------------------------------------------------------ //
type ReportBuilder struct {
	DefaultAdapter
}

// Fonction GetParser de l'objet ReportBuilder
func (r ReportBuilder) GetParser() *tools.HTMLParser {
	r.parser = &tools.HTMLParser{"Report", nil, nil}
	return r.parser
}
