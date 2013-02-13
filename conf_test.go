package conf

import (
	"fmt"
	"testing"
)

const confFile = `
[default]
host = example.com
port = 43
compression = on
active = false
float = 2.3

[service-1]
port = 443

[list]
list-1 =
	one, one
	two
	three
list-2 = "one,one", two, three
list-3 = one,one
	two

list-4 = 1, 2, 3
list-5 = 1.3, 2.0, 3

list-6 = yes, true, no, 0, n
`

//url = http://%(host)s/something

type stringTest struct {
	section string
	option  string
	answer  string
}
type stringListTest struct {
	section string
	option  string
	answer  []string
}

type intTest struct {
	section string
	option  string
	answer  int
}
type intListTest struct {
	section string
	option  string
	answer  []int
}

type int64Test struct {
	section string
	option  string
	answer  int64
}
type int64ListTest struct {
	section string
	option  string
	answer  []int64
}

type float64Test struct {
	section string
	option  string
	answer  float64
}
type float64ListTest struct {
	section string
	option  string
	answer  []float64
}

type boolTest struct {
	section string
	option  string
	answer  bool
}
type boolListTest struct {
	section string
	option  string
	answer  []bool
}

var testSet = []interface{}{
	stringTest{"", "host", "example.com"},
	stringListTest{"list", "list-1", []string{"one, one", "two", "three"}},
	intTest{"default", "port", 43},
	intListTest{"list", "list-4", []int{1, 2, 3}},
	int64Test{"default", "port", 43},
	int64ListTest{"list", "list-4", []int64{1, 2, 3}},
	float64Test{"default", "float", 2.3},
	float64ListTest{"list", "list-5", []float64{1.3, 2.0, 3}},
	boolTest{"default", "compression", true},
	boolTest{"default", "active", false},
	boolListTest{"list", "list-6", []bool{true, true, false, false, false}},
	intTest{"service-1", "port", 443},
	//stringTest{"service-1", "url", "http://example.com/something"},
}

func TestBuild(t *testing.T) {
	c, err := ReadBytes([]byte(confFile))
	if err != nil {
		t.Error(err)
	}

	for testnum, element := range testSet {
		switch element.(type) {
		case stringTest:
			e := element.(stringTest)
			ans, err := c.String(e.section, e.option)
			verify(t, testnum, "c.String", e.section, e.option, ans, e.answer, err)
		case stringListTest:
			e := element.(stringListTest)
			ans, err := c.StringList(e.section, e.option)
			verifyList(t, testnum, "c.StringList", e.section, e.option, ans, e.answer, err)
		case intTest:
			e := element.(intTest)
			ans, err := c.Int(e.section, e.option)
			verify(t, testnum, "c.Int", e.section, e.option, ans, e.answer, err)
		case intListTest:
			e := element.(intListTest)
			ans, err := c.IntList(e.section, e.option)
			verifyList(t, testnum, "c.IntList", e.section, e.option, ans, e.answer, err)
		case int64Test:
			e := element.(int64Test)
			ans, err := c.Int64(e.section, e.option)
			verify(t, testnum, "c.Int64", e.section, e.option, ans, e.answer, err)
		case int64ListTest:
			e := element.(int64ListTest)
			ans, err := c.Int64List(e.section, e.option)
			verifyList(t, testnum, "c.Int64List", e.section, e.option, ans, e.answer, err)
		case float64Test:
			e := element.(float64Test)
			ans, err := c.Float64(e.section, e.option)
			verify(t, testnum, "c.Float64", e.section, e.option, ans, e.answer, err)
		case float64ListTest:
			e := element.(float64ListTest)
			ans, err := c.Float64List(e.section, e.option)
			verifyList(t, testnum, "c.Float64List", e.section, e.option, ans, e.answer, err)
		case boolTest:
			e := element.(boolTest)
			ans, err := c.Bool(e.section, e.option)
			verify(t, testnum, "c.Bool", e.section, e.option, ans, e.answer, err)
		case boolListTest:
			e := element.(boolListTest)
			ans, err := c.BoolList(e.section, e.option)
			verifyList(t, testnum, "c.BoolList", e.section, e.option, ans, e.answer, err)
		}
	}
}

func verify(t *testing.T, testnum int, testcase, section, option string, output, expected interface{}, err error) {
	if err != nil {
		t.Fatalf(`%d. %s("%s", "%s") returned error: %v`, testnum, testcase, section, option, err.Error())
	}
	if output != expected {
		t.Fatalf(`%d. %s("%s", "%s"): output %v != %v`, testnum, testcase, section, option, output, expected)
	}
}

func verifyList(t *testing.T, testnum int, testcase, section, option string, output, expected interface{}, err error) {
	if err != nil {
		t.Fatalf(`%d. %s("%s", "%s") returned error: %v`, testnum, testcase, section, option, err.Error())
	}
	outputF := fmt.Sprintf("%v", output)
	expectedF := fmt.Sprintf("%v", expected)
	if outputF != expectedF {
		t.Fatalf(`%d. %s("%s", "%s"): output %v != %v`, testnum, testcase, section, option, output, expected)
	}
}
