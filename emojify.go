package emojify

import (
	"strings"
)

func Emojify(s string) string {
	for _, shortcode := range EMOJI_SHORT_CODES {
		s = replaceWithShortcode(s, shortcode)
	}

	return s
}

func replaceWithShortcode(s string, shortcode string) string {
	startIndex := strings.Index(s, shortcode)

	if startIndex == -1 {
		return s
	}

	endIndex := (startIndex) + len(shortcode)

	if !isInsideEmoji(s, startIndex, endIndex) {
		return s[0:startIndex] + ":" + shortcode + ":" + replaceWithShortcode(s[endIndex:], shortcode)
	} else {
		return s[0:endIndex] + replaceWithShortcode(s[endIndex:], shortcode)
	}

}

func isInsideEmoji(s string, startIndex int, endIndex int) bool {
	endingMatches := 0
	stringLen := len(s)

	for {
		if endIndex > stringLen-1 {
			break
		}
		endChar := s[endIndex]

		if endChar == ':' {
			endingMatches += 1
		} else if endChar == ' ' || endChar == ',' || endChar == ';' {
			break
		}
		endIndex += 1
	}

	if endingMatches == 0 {
		return false
	} else {
		return endingMatches%2 == 1
	}
}
