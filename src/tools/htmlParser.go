package tools

import (
	"exp/html"
	"fmt"
	"io"
	"log"
	"strings"

// "studash/adapters"
// "io/ioutil"
// "net/http"
// "time"
)

type HTMLParameter struct {
	Name string
	Value string
}

type NodeList []*html.Node

type HTMLParser struct {
	Msg string
	Data io.Reader
	HTMLNodes []html.Node
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
func (p *HTMLParser) GetValue(tag string, feilds ...HTMLParameter) []html.Node {
	aNode, errors := html.Parse(p.Data)
	
	if errors != nil {
		log.Println("ERROR : error while parsing web page")
		return nil
	}
	
	// recursive search
	p.GetNodes(aNode, tag, feilds)
	
	return p.HTMLNodes
}

func (p *HTMLParser) GetNodes(aNode *html.Node, tag string, feilds []HTMLParameter) {
	
	// check for the tag we specified
	if (*aNode).Data == tag || tag == ""{
		legit := true
		
		// We get threw all attributes of the tag
		for _, attr := range (*aNode).Attr {
			
			// look if the tag meets all the requirements
			for _, feild := range feilds {
				// if one of the attributes doesn't have the good value
				// then the node isn't interesting anymore
				if feild.Name == attr.Key && (feild.Value != attr.Val || feild.Value != "") {
					legit = false
					break
				}
			}
			if !legit {
				break
			}
		}
		// if the node is legit we append it to the list
		if legit {
			p.HTMLNodes = append(p.HTMLNodes, *aNode)
		}
	}
	//
	for _, nodey := range (*aNode).Child {
		p.GetNodes(nodey, tag, feilds)
	}
}

func FineSearch( aNode html.Node, wanted string, tag string, feilds ...HTMLParameter) (bool, string){
	legit := true
	value := ""
	
	if aNode.Data == tag || tag == ""{
		
		// We get threw all attributes of the tag
		for _, attr := range aNode.Attr {
			
			// look if the tag meets all the requirements
			for _, feild := range feilds {
				if attr.Key == wanted {
					value = attr.Val
				}
				// if one of the attributes doesn't have the good value
				// then the node isn't interesting anymore
				// (!strings.Contains(attr.Val, feild.Value))
				if (feild.Name == attr.Key) && (!strings.Contains(attr.Val, feild.Value)) {
					legit = false
					break
				} 
			}
			if !legit {
				value = ""
				break
			}
		}
	} else {
		legit = false
	}
	return legit, value
}

func WeirdJSON() []byte {
	return []byte("{name:lolsaure}")
	// return pages.KAR
}
