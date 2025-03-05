package prose

import "strings"

func JoinWithComas(phrases []string) string {
	l := len(phrases)

	if l == 0 {
		return ""
	} else if l == 2 {
		return phrases[0] + " and " + phrases[1]
	} else if l == 1 {
		return phrases[0]
	} else {
		result := strings.Join(phrases[:l-1], ", ")
		result += ", and "
		result += phrases[l-1]
		return result
	}
}
