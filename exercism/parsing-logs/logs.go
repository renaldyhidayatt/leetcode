package parsinglogs

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-=~*]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)

	i := 0

	for _, line := range lines {
		if re.MatchString(line) {
			i++
		}
	}

	return i
}

func RemoveEndOfLineText(text string) string {

	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUsername(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)

	s := []string{}

	for _, line := range lines {
		founds := re.FindStringSubmatch(line)

		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}

		s = append(s, line)
	}

	return s
}
