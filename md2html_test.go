package md2html_test

import (
	"testing"

	"github.com/MarkyMan4/md2html"
)

func TestConvert(t *testing.T) {
	md2html.Convert("# hello")
}
