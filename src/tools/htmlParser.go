package tools

import (
	"exp/html"
	"fmt"
	"io"
	"log"

// "studash/adapters"
// "io/ioutil"
// "net/http"
// "time"
)

type HTMLParameter struct {
	Name, Value string
}

type NodeList []*html.Node

type HTMLParser struct {
	Msg       string
	Data      io.Reader
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
	if (*aNode).Data == tag || tag == "" {
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

type AttributeSet []html.Attribute

func (a AttributeSet) Contains(rhs html.Attribute) bool {
	for _, attribute := range a {
		if attribute.Namespace == rhs.Namespace &&
			attribute.Key == rhs.Key &&
			attribute.Val == rhs.Val {
			return true
		}
	}
	return false
}

func (a AttributeSet) Key(key string) AttributeSet {
	var buffer AttributeSet

	for _, attribute := range a {
		if attribute.Key == key {
			buffer = append(buffer, attribute)
		}
	}
	return buffer
}

func (a AttributeSet) First() html.Attribute {
	if len(a) == 0 {
		return html.Attribute{"undefined", "undefined", "undefined"}
	}
	return a[0]
}
