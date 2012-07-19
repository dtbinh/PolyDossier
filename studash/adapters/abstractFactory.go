package adapters

import (
	"studash/tools"
)

// ------------------------------------------------------ //
//                    Builder d'adapters                  //
// ------------------------------------------------------ //
type AdapterBuilder interface {
	GetParser() *tools.HTMLParser
}

// ------------------------------------------------------ //
//              Builder pour la page de login             //
// ------------------------------------------------------ //
type LoginBuilder struct {
	parser *tools.HTMLParser
}
// Fonction GetParser de l'objet LoginBuilder
func (l LoginBuilder) GetParser() *tools.HTMLParser {
	l.parser = new(tools.HTMLParser)
	return l.parser
}

// ------------------------------------------------------ //
//                  Fonction Main de test                 //
// ------------------------------------------------------ //
// func main() {
	// var ab AdapterBuilder
	// parser := ab.GetParser()
	// parser.Print()
// }
