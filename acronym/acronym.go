package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	letters := regexp.MustCompile("\\b\\w").FindAllString(s, -1)
	return strings.ToUpper(strings.Join(letters, ""))
}
