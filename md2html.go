package md2html

import (
	"encoding/xml"
	"regexp"
	"strings"
)

var (
	h1Regex = regexp.MustCompile("^# ")
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

type HtmlBody struct {
	Elements []*HtmlElement
}

func NewHtmlBody() *HtmlBody {
	return &HtmlBody{
		Elements: []*HtmlElement{},
	}
}

func (hb *HtmlBody) Render() string {
	res, err := xml.MarshalIndent(hb.Elements, "", "    ")
	if err != nil {
		panic("failed to render HTML document " + err.Error())
	}

	return string(res)
}

func ParseMarkdown(md string) *HtmlBody {
	html := NewHtmlBody()

	lines := strings.Split(md, "\n")
	for _, line := range lines {
		if matched := h1Regex.Match([]byte(line)); matched {
			h1 := NewHtmlElement("h1", line, "")
			html.Elements = append(html.Elements, h1)
		}
	}

	return html
}
