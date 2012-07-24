package adapters

import (
	"studash/tools"
)

// func (*DefaultAdapter) Run (r *http.Request) {
// resp, err := goResponse(r)
// parser = make.HTMLParser(resp)

// // TODO: put here what we exactly need
// the_first_value := parser.getFeild("MAP OF FEILDS HERE")
// the_second_value := parser.getFeild("MAP OF FEILDS HERE")
// }

// ------------------------------------------------------ //
//              Builder pour la page de login             //
// ------------------------------------------------------ //
type LoginBuilder struct {
	DefaultAdapter
}

// Fonction GetParser de l'objet LoginBuilder
func (l LoginBuilder) GetParser() *tools.HTMLParser {
	l.parser = &tools.HTMLParser{"Login"}
	return l.parser
}
