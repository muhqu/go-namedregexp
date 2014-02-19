package namedregexp

import (
	"fmt"
	"testing"
)

func TestFindNamedStringSubmatch(t *testing.T) {
	var myExp = MustCompile(`^((?P<last>\w+),\s*(?P<first>\w+)(\s+(?P<middle>\w+))?|(?P<first>\w+)\s+((?P<middle>\w+)\s+)?(?P<last>\w+))$`)
	tests := []string{
		"John||Doe", "John Doe",
		"Foo||Bar", "Bar, Foo",
		"Muh||Cow", "Cow,Muh",
		"Richard|Dean|Anderson", "Anderson, Richard Dean",
	}
	for i := 0; i < len(tests); i += 2 {
		var testInput = tests[i+1]
		var expectedOutput = tests[i]
		var matches = myExp.FindNamedStringSubmatch(testInput)
		var actualOutput = fmt.Sprintf("%s|%s|%s", matches["first"], matches["middle"], matches["last"])
		if actualOutput != expectedOutput {
			t.Errorf("for input '%s' the output '%s' doesn't match expected '%s'\n", testInput, actualOutput, expectedOutput)
		}
	}
}

func TestFindNamedStringSubmatchIndex(t *testing.T) {
	var myExp = MustCompile(`^((?P<last>\w+),\s*(?P<first>\w+)(\s+(?P<middle>\w+))?|(?P<first>\w+)\s+((?P<middle>\w+)\s+)?(?P<last>\w+))$`)
	tests := map[string][][]int{
		"John Doe":               {{0, 4}, nil, {5, 8}},
		"Bar, Foo":               {{5, 8}, nil, {0, 3}},
		"Cow,Muh":                {{4, 7}, nil, {0, 3}},
		"Anderson, Richard Dean": {{10, 17}, {18, 22}, {0, 8}},
	}
	// t.Errorf("%+v\n", tests)
	for testInput, expectedMatches := range tests {
		m := myExp.FindNamedStringSubmatchIndex(testInput)
		actualMatches := ArArInt{m["first"], m["middle"], m["last"]}

		ok, err := actualMatches.equals(expectedMatches)
		if !ok {
			t.Errorf("for input '%s': %s\n", testInput, err)
			t.FailNow()
		}
	}
}

type ArArInt [][]int

func (expected ArArInt) equals(actual ArArInt) (bool, error) {
	if len(expected) != len(actual) {
		return false, fmt.Errorf("invalid length")
	}
	for i, _ := range expected {
		if len(expected[i]) != len(actual[i]) {
			return false, fmt.Errorf("invalid length at index [%d]", i)
		}
		for k, _ := range expected[i] {
			if actual[i][k] != expected[i][k] {
				return false, fmt.Errorf("%d != %d at index [%d][%d]", actual[i][k], expected[i][k], i, k)
			}
		}
	}
	return true, nil
}
