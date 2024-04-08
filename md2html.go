package md2html

import (
	"encoding/xml"
	"regexp"
	"strings"
)

type HtmlElement struct {
	XMLName xml.Name
	Id      string `xml:"id,attr,omitempty"`
	Content string `xml:",chardata"`
}

func NewHtmlElement(tagName string, content string, id string) *HtmlElement {
	elem := &HtmlElement{
		XMLName: xml.Name{Local: tagName},
		Content: content,
		Id:      id,
	}

	return elem
}

type HtmlDoc struct {
	XMLName xml.Name `xml:"html"`
	Head    *HtmlHead
	Body    *HtmlBody
}

type HtmlHead struct {
	XMLName xml.Name `xml:"head"`
	Title   string   `xml:"title"`
}
type HtmlBody struct {
	XMLName  xml.Name `xml:"body"`
	Elements []*HtmlElement
}

func NewHtmlDoc(title string) *HtmlDoc {
	return &HtmlDoc{
		Head: &HtmlHead{Title: title},
		Body: &HtmlBody{Elements: []*HtmlElement{}},
	}
}

func (hd *HtmlDoc) Render() string {
	rendered := "<!DOCTYPE html>\n"

	res, err := xml.MarshalIndent(hd, "", "    ")
	if err != nil {
		panic("failed to render HTML document " + err.Error())
	}

	return rendered + string(res)
}

func Convert(md string) string {
	htmlDoc := NewHtmlDoc("Markdown")

	lines := strings.Split(md, "\n")
	for _, line := range lines {
		if matched, _ := regexp.Match("^# ", []byte(line)); matched {
			h1 := NewHtmlElement("h1", line, "")
			htmlDoc.Body.Elements = append(htmlDoc.Body.Elements, h1)
		}
	}

	return htmlDoc.Render()
}
