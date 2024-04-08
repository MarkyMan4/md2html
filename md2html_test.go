package md2html_test

import (
	"fmt"
	"testing"

	"github.com/MarkyMan4/md2html"
)

const md = `# hello
# this is a test
`

func TestConvert(t *testing.T) {
	res := md2html.ParseMarkdown(md)
	fmt.Println(res.Render())
}
