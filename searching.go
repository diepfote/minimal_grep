package main

import (
	"regexp"
	"strings"
)

func searchString(pattern string, content string, ignoreCase bool) []string {
	if ignoreCase {
		pattern = strings.ToLower(pattern)
		content = strings.ToLower(content)
	}

	lines := strings.Split(content, "\n")
	var matches []string

	linenumber := 1
	for _, line := range lines {
		if strings.Contains(line, pattern) {
			matches = append(matches, line+"\n")
		}
		linenumber++
	}

	return matches
}

func searchRegex(pattern string, content string, ignoreCase bool) []string {
	if ignoreCase {
		//panic("Not implemented")
		//pattern = "(?i)" + pattern
		pattern = strings.ToLower(pattern)
		content = strings.ToLower(content)
	}

	re := regexp.MustCompilePOSIX(pattern + ".*[\n]")
	return re.FindAllString(content, -1)
}
