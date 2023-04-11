package util

import (
	"github.com/microcosm-cc/bluemonday"
	"regexp"
)

var p = bluemonday.StripTagsPolicy()
var regex = regexp.MustCompile("\n[ \n]{2,}")

func FormatRSSContent(content string) string {
	return regex.ReplaceAllString(p.Sanitize(content), "\n")
}
