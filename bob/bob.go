package bob

import "strings"

const testVersion = 3

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	if strings.ToLower(remark) != remark && strings.ToUpper(remark) == remark {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(remark, "?") {
		return "Sure."
	}
	return "Whatever."
}
