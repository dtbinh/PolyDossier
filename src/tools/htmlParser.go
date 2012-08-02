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

////
// To get the value of a feild in the HTML
// 
//  string : the name of the attribute that we want to get the value of
//  string : the name of the tag we want to search in
//  HTMLParameter : pairs of name/value we want in the balise
// 
//  return : the value of the wanted parameter
////
func (p *HTMLParser) GetValue(wanted , tag string, champs ...HTMLParameter) interface{} {
	aNode, errors := html.Parse(p.Data)
	if errors != nil {
		fmt.Println("Error :", errors)
		return nil
	}
	
	// recursive search
	result := printNodeContent(aNode, wanted, tag, champs)
	fmt.Println("Result :", result)
	
	return nil
}

// function de test
func printNodeContent(aNode *html.Node, wanted string, tag string, champs []HTMLParameter) string {
	
	// if the tag has no child it means it's a leaf
	// and we want to check what's inside
	if aNode.Child == nil {
		if (*aNode).Data == tag {
			x := ""
			legit := true
			
			// We get threw all attributes of the tag
			for _, attr := range (*aNode).Attr {
				if attr.Key == wanted {
					x = attr.Val
				}
				
				// look if the tag meets all the requirements
				for _, champ := range champs {
					// if one of the attributes doesn't have the goo value
					// then the node isn't interesting anymore
					if champ.Name == attr.Key && champ.Value != attr.Val {
						legit = false
						break
					}
				}
				if !legit {
					break
				}
			}
			
			// if the tag is legit and we found a value
			// we must stop the search and return the result
			if legit && x != "" {
				return x
			}
		}
	} else {
		// we get here only if the tag has childrens
		for _, nodey := range (*aNode).Child {
			result := printNodeContent(nodey, wanted, tag, champs)
			// we don't neet to continue the search if we found
			// what we are looking for
			if result != "" {
				return result
			}
		}
	}
		
	// if we found nothing... well, we return nothing :O
	return ""
}

func WeirdJSON() []byte {
	return []byte("{name:lolsaure}")
	// return pages.KAR
}
