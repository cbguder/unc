package evernote

import "encoding/xml"

type enexExport struct {
	XMLName string     `xml:"en-export"`
	Notes   []enexNote `xml:"note"`
}

type enexNote struct {
	Title    string      `xml:"title"`
	Content  enexContent `xml:"content"`
	Created  string      `xml:"created"`
	Modified string      `xml:"updated"`
	Tags     []string    `xml:"tag"`
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
	XMLName  xml.Name
	CharData string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}
