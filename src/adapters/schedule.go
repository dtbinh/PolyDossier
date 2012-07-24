package adapters

import (
	"studash/tools"
)

// ------------------------------------------------------ //
//               Builder pour la page Horaire             //
// ------------------------------------------------------ //
type ScheduleBuilder struct {
	DefaultAdapter
}

// Fonction GetParser de l'objet ScheduleBuilder
func (s ScheduleBuilder) GetParser() *tools.HTMLParser {
	s.parser = &tools.HTMLParser{"Schedule"}
	return s.parser
}
