package on99

import "strings"

func StrStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
