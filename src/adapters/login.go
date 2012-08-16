package adapters

import (
	"exp/html"
	"io"
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
func (l *LoginBuilder) SetParser(r io.Reader) {
	if l.parser == nil {
		l.parser = &tools.HTMLParser{"Login", r, nil}
	}
}

func (l *LoginBuilder) GetMatriculeToken(r io.Reader) (string, string) {
	l.SetParser(r)
	nodes := l.parser.GetValue("input")
	var matricule, token string
	for _, node := range nodes {

		// Si on a les deux
		if len(matricule) > 0 && len(token) > 0 {
			return matricule, token
		}

		if len(matricule) == 0 && tools.AttributeSet(node.Attr).Contains(html.Attribute{"", "name", "matricule"}) {
			matricule = tools.AttributeSet(node.Attr).Key("value").First().Val
		}

		if len(token) == 0 && tools.AttributeSet(node.Attr).Contains(html.Attribute{"", "name", "token"}) {
			token = tools.AttributeSet(node.Attr).Key("value").First().Val
		}
	}
	return matricule, token
}
