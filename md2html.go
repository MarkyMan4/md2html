package md2html

import (
	"encoding/xml"
	"regexp"
	"strings"
)

var (
	h1Regex = regexp.MustCompile("^# (.*)$")
	h2Regex = regexp.MustCompile("^## (.*)$")
	h3Regex = regexp.MustCompile("^### (.*)$")
	h4Regex = regexp.MustCompile("^#### (.*)$")
	h5Regex = regexp.MustCompile("^##### (.*)$")
	h6Regex = regexp.MustCompile("^###### (.*)$")
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

func getHeaderElement(rawText string, regex *regexp.Regexp, tag string) *HtmlElement {
	innerText := string(regex.ReplaceAll([]byte(rawText), []byte("$1")))
	element := NewHtmlElement(tag, innerText, "")

	return element
}

func ParseMarkdown(md string) *HtmlBody {
	html := NewHtmlBody()

	lines := strings.Split(md, "\n")
	for _, line := range lines {
		var element *HtmlElement

		if matched := h1Regex.Match([]byte(line)); matched {
			element = getHeaderElement(line, h1Regex, "h1")
		} else if matched := h2Regex.Match([]byte(line)); matched {
			element = getHeaderElement(line, h2Regex, "h2")
		} else if matched := h3Regex.Match([]byte(line)); matched {
			element = getHeaderElement(line, h3Regex, "h3")
		} else if matched := h4Regex.Match([]byte(line)); matched {
			element = getHeaderElement(line, h4Regex, "h4")
		} else if matched := h5Regex.Match([]byte(line)); matched {
			element = getHeaderElement(line, h5Regex, "h5")
		} else if matched := h6Regex.Match([]byte(line)); matched {
			element = getHeaderElement(line, h6Regex, "h6")
		}

		if element != nil {
			html.Elements = append(html.Elements, element)
		}
	}

	return html
}
