package md2html

import (
	"fmt"
	"strings"
)

func Convert(md string) string {
	lines := strings.Split(md, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	return ""
}
