package md2html_test

import (
	"fmt"
	"testing"

	"github.com/MarkyMan4/md2html"
)

func TestConvert(t *testing.T) {
	res := md2html.Convert("# hello")
	fmt.Println(res)
}
