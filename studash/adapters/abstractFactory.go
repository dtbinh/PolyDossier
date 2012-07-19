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
