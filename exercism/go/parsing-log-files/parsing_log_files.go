package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*$`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			count += len(matches)
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	userLogRegexp := regexp.MustCompile(`User\s+(\w+)`)
	nameRegexp := regexp.MustCompile(`[\w\d]+$`)

	for index, line := range lines {
		logMatch := userLogRegexp.FindString(line)
		if logMatch != "" {
			userName := nameRegexp.FindString(logMatch)
			log := fmt.Sprintf("[USR] %s %s", userName, line)
			lines[index] = log
		}
	}

	return lines
}
