package main

import (
	//"io"
//	"text/template"
//	html_template "html/template"
	"io"
)




type Node struct{
	Type     NodeType
	Data  	string
	Attr     []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32
const(
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommonNode
	DoctypeNode
)
type Attribute struct {
	Key, Val string
}

//Parse
func Parse(r io.Reader) (*Node, error) {
}
