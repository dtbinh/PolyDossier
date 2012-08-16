package adapters

import (
	"studash/tools"
	"io"
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
func (l *LoginBuilder) SetParser(r io.Reader) {
	if(l.parser == nil) {
		l.parser = &tools.HTMLParser{"Login", r, nil}
	}
}

func (l *LoginBuilder) GetMatriculeToken(r io.Reader) (string,string) {
	l.SetParser(r)
	nodes := l.parser.GetValue("input")
	var matricule, token string
	for _, node := range nodes {
		b1, v1 := tools.FineSearch(node, "value", "", tools.HTMLParameter{"name", "matricule"})
		b2, v2 := tools.FineSearch(node, "value", "", tools.HTMLParameter{"name", "token"})
		
		if b1 {
			matricule = v1
		}
		if b2 {
			token = v2
		}
		if b1 && b2 {
			break
		}
	}
	return matricule, token
}
