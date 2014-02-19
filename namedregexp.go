package namedregexp

import (
	"regexp"
)

type NamedRegexp struct {
	*regexp.Regexp
}
type NamedStringSubmatch map[string]string
type NamedStringSubmatchIndex map[string][]int

func MustCompile(str string) *NamedRegexp {
	var namedExp = new(NamedRegexp)
	namedExp.Regexp = regexp.MustCompile(str)
	return namedExp
}

func (namedExp *NamedRegexp) FindNamedStringSubmatch(str string) NamedStringSubmatch {
	return FindNamedStringSubmatch(namedExp.Regexp, str)
}

func (namedExp *NamedRegexp) FindNamedStringSubmatchIndex(str string) NamedStringSubmatchIndex {
	return FindNamedStringSubmatchIndex(namedExp.Regexp, str)
}

func FindNamedStringSubmatch(exp *regexp.Regexp, str string) NamedStringSubmatch {
	names := exp.SubexpNames()
	matches := exp.FindStringSubmatch(str)
	if matches == nil {
		return nil
	}
	result := make(NamedStringSubmatch, len(names))
	for i, name := range names {
		if name != "" {
			_, present := result[name]
			if matches[i] != "" || !present {
				result[name] = matches[i]
			}
		}
	}
	return result
}

func FindNamedStringSubmatchIndex(exp *regexp.Regexp, str string) NamedStringSubmatchIndex {
	names := exp.SubexpNames()
	matches := exp.FindStringSubmatchIndex(str)
	if matches == nil {
		return nil
	}
	result := make(NamedStringSubmatchIndex, len(names))
	for i, name := range names {
		if name != "" {
			_, present := result[name]
			match := matches[(i * 2):(i*2 + 2)]
			if match[0] != -1 {
				result[name] = match
			} else if !present {
				result[name] = nil
			}
		}
	}
	return result
}
