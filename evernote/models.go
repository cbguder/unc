package evernote

import "encoding/xml"

type enexExport struct {
	XMLName string     `xml:"en-export"`
	Notes   []enexNote `xml:"note"`
}

type enexNote struct {
	Title    string      `xml:"title"`
	Content  enexContent `xml:"content"`
	Tags     []string    `xml:"tag"`
	Created  string      `xml:"created"`
	Modified string      `xml:"updated"`
}

type enexContent struct {
	XMLName string `xml:"content"`
	Body    string `xml:",cdata"`
}

type enexInnerNote struct {
	XMLName  string `xml:"en-note"`
	Children []enexNode
}

type enexNode struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}
