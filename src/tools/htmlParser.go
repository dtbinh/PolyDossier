package tools

import (
	"exp/html"
	"fmt"
	"io"

// "studash/adapters"
// "io/ioutil"
// "net/http"
// "time"
)

type HTMLParameter struct {
	Name string
	Value string
}

type HTMLParser struct {
	Msg string
	Data io.Reader
}

func (p *HTMLParser) Print() {
	fmt.Println(p.Msg)
}

// permet de récupérer la valeur d'un champ dans une balise HTML
// le type retourné est inconnu car il peut autant s'agir d'une 
// string que d'un nombre

///
// To get the value of a feild in the HTML
// 
//  string : the name of the attribute that we want to get the value of
//  string : the name of the balise we want to search in
//  HTMLParameter : pairs of name/value we want in the balise
// 
//  return : the value of the wanted parameter
///
func (p *HTMLParser) GetValue(wanted string, balise string, champs ...HTMLParameter) interface{} {
	aNode, errors := html.Parse(p.Data)
	if errors != nil {
		fmt.Println("Error :", errors)
		return nil
	}
	
	// recursive search
	printNodeContent(aNode, wanted, balise)
	
	return nil
}

// function de test
func printNodeContent(aNode *html.Node, wanted string, balise string) {
	if aNode.Child == nil {
		fmt.Println(*aNode)
		return
	}
	
	for _, nodey := range (*aNode).Child {
		printNodeContent(nodey, wanted, balise)
	}
}

// func ParseHTML(w http.ResponseWriter, r *http.Request) []byte {
// t := time.Now()
// resp, err := GoResponse(r)
// respCpy, err := GoResponse(r)

// if err != nil {
// fmt.Println(r.Method, err)
// return WeirdJSON()
// }

// z := html.NewTokenizer(resp.Body)

// for {
// tt := z.Next()
// if tt == html.ErrorToken {
// fmt.Println("Error token")
// break
// }
// }

// data, err := ioutil.ReadAll(respCpy.Body)
// if err != nil {
// fmt.Println(r.Method, err)
// return WeirdJSON()
// }

// resp.Body.Close()
// respCpy.Body.Close()

// elapsed := time.Since(t)

// if elapsed < 200*time.Millisecond {
// fmt.Println(time.Since(t).String() + "\t" + r.URL.Path)
// } else {
// fmt.Println(time.Since(t).String() + "\t" + r.URL.Path + "\t[Trop Long!!!]")
// }

// return data
// }

func WeirdJSON() []byte {
	return []byte("{name:lolsaure}")
	// return pages.KAR
}
