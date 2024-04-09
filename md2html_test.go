package md2html_test

import (
	"fmt"
	"testing"

	"github.com/MarkyMan4/md2html"
)

const md = `# heading 1
## heading 2
### heading 3
#### heading 4
##### heading 5
###### heading 6
`

func TestConvert(t *testing.T) {
	res := md2html.ParseMarkdown(md)
	fmt.Println(res.Render())

	if len(res.Elements) != 6 {
		t.Fatalf("Incorrect number of elements parsed, expected 6, got %d\n", len(res.Elements))
	}

	// TODO check tags that were generated
}
