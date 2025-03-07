package pdf

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// example input: "1,3-5,7"
// example output: []string{"1", "3", "4", "5", "7"}
func ParsePageRange(pageRange string) ([]string, error) {
	if pageRange == "" {
		return nil, errors.New("page range cannot be empty")
	}

	var pages []string
	segments := strings.Split(pageRange, ",")
	rangePattern := regexp.MustCompile(`^\d+-\d+$`)
	singlePagePattern := regexp.MustCompile(`^\d+$`)

	for _, segment := range segments {
		segment = strings.TrimSpace(segment)
		if rangePattern.MatchString(segment) {
			parts := strings.Split(segment, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			if start > end {
				return nil, errors.New("invalid range: start page is greater than end page")
			}
			for i := start; i <= end; i++ {
				pages = append(pages, strconv.Itoa(i))
			}
		} else if singlePagePattern.MatchString(segment) {
			pages = append(pages, segment)
		} else {
			return nil, errors.New("invalid page range format")
		}
	}

	return pages, nil
}
