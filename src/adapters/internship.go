package adapters

import (
	"studash/tools"
)

// ------------------------------------------------------ //
//                Builder pour la page Stage              //
// ------------------------------------------------------ //
type InternshipBuilder struct {
	DefaultAdapter
}

// Fonction GetParser de l'objet InternshipBuilder
func (i InternshipBuilder) GetParser() *tools.HTMLParser {
	i.parser = &tools.HTMLParser{"Internship", nil, nil}
	return i.parser
}
