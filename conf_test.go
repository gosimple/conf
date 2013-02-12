package conf

import (
	"fmt"
	"strconv"
	"testing"
)

const confFile = `
[default]
host = example.com
port = 43
compression = on
active = false

[service-1]
port = 443
`

//url = http://%(host)s/something

type stringtest struct {
	section string
	option  string
	answer  string
}

type inttest struct {
	section string
	option  string
	answer  int
}

type int64test struct {
	section string
	option  string
	answer  int64
}

type booltest struct {
	section string
	option  string
	answer  bool
}

var testSet = []interface{}{
	stringtest{"", "host", "example.com"},
	inttest{"default", "port", 43},
	int64test{"default", "port", 43},
	booltest{"default", "compression", true},
	booltest{"default", "active", false},
	inttest{"service-1", "port", 443},
	//stringtest{"service-1", "url", "http://example.com/something"},
}

func TestBuild(t *testing.T) {
	c, err := ReadBytes([]byte(confFile))
	if err != nil {
		t.Error(err)
	}

	for _, element := range testSet {
		switch element.(type) {
		case stringtest:
			e := element.(stringtest)
			ans, err := c.String(e.section, e.option)
			if err != nil {
				t.Error("c.String(\"" + e.section + "\",\"" + e.option + "\") returned error: " + err.Error())
			} else if ans != e.answer {
				t.Error("c.String(\"" + e.section + "\",\"" + e.option + "\") returned incorrect answer: " + ans)
			}
		case inttest:
			e := element.(inttest)
			ans, err := c.Int(e.section, e.option)
			if err != nil {
				t.Error("c.Int(\"" + e.section + "\",\"" + e.option + "\") returned error: " + err.Error())
			} else if ans != e.answer {
				t.Error("c.Int(\"" + e.section + "\",\"" + e.option + "\") returned incorrect answer: " + strconv.Itoa(ans))
			}
		case int64test:
			e := element.(int64test)
			ans, err := c.Int64(e.section, e.option)
			if err != nil {
				t.Error("c.Int64(\"" + e.section + "\",\"" + e.option + "\") returned error: " + err.Error())
			} else if ans != e.answer {
				ans64 := fmt.Sprintf("%v", ans)
				t.Error("c.Int64(\"" + e.section + "\",\"" + e.option + "\") returned incorrect answer: " + ans64)
			}
		case booltest:
			e := element.(booltest)
			ans, err := c.Bool(e.section, e.option)
			if err != nil {
				t.Error("c.Bool(\"" + e.section + "\",\"" + e.option + "\") returned error: " + err.Error())
			} else if ans != e.answer {
				t.Error("c.Bool(\"" + e.section + "\",\"" + e.option + "\") returned incorrect answer")
			}
		}
	}
}
