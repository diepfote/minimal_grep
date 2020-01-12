package main

import (
	"regexp"
	"strings"
)

func searchString(pattern string, content string, ignoreCase bool, lineByLine bool) map[int]string {
	if ignoreCase {
		pattern = strings.ToLower(pattern)
		content = strings.ToLower(content)
	}

	lines := strings.Split(content, "\n")
	matches := make(map[int]string)

	for index, line := range lines {
		if strings.Contains(line, pattern) {
			matches[index] = line + "\n"
		}
	}

	return matches
}

func searchRegex(pattern string, content string, ignoreCase bool, lineByLine bool) map[int]string {
	if ignoreCase {
		//panic("Not implemented")
		//pattern = `(?i)` + pattern
		pattern = strings.ToLower(pattern)
		content = strings.ToLower(content)
	}

	re := regexp.MustCompilePOSIX(pattern + ".*[\n]{0,1}")
	matches := re.FindAllString(content, -1)

	results := make(map[int]string)
	if !lineByLine {
		for index, match := range matches {
			results[index] = match
		}
	} else {
		lines := strings.Split(content, "\n")
		for index, line := range lines {
			match := re.FindString(line)
			if len(match) > 0 {
				results[index] = match + "\n"
			}
		}
	}

	return results
}
