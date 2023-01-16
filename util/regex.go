package util

import "regexp"

// takes a regular expression and a slice of strings
// the return value is a slice containing, for each string, if there's a match,
// the match and the matches of any subexpressions.
func FindStringSubmatchesInSlice(lines []string, regexString string) [][]string {
	output := [][]string{}
	r := regexp.MustCompile(regexString)
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		if len(m) > 0 {
			output = append(output, m)
		}
	}

	return output
}
