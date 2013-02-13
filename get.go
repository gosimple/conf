package conf

import (
	"encoding/csv"
	"strconv"
	"strings"
)

// Sections returns the list of sections in the configuration.
// (The default section always exists.)
func (c *Config) Sections() (sections []string) {
	sections = make([]string, len(c.data))

	i := 0
	for s, _ := range c.data {
		sections[i] = s
		i++
	}

	return sections
}

// HasSection checks if the configuration has the given section.
// (The default section always exists.)
func (c *Config) HasSection(section string) bool {
	if section == "" {
		section = "default"
	}
	_, ok := c.data[strings.ToLower(section)]

	return ok
}

// Options returns the list of options available in the given section.
// It returns an error if the section does not exist and an empty list if the section is empty.
// Options within the default section are also included.
func (c *Config) Options(section string) (options []string, err error) {
	if section == "" {
		section = "default"
	}
	section = strings.ToLower(section)

	if _, ok := c.data[section]; !ok {
		return nil, GetError{SectionNotFound, "", "", section, ""}
	}

	options = make([]string, len(c.data[DefaultSection])+len(c.data[section]))
	i := 0
	for s, _ := range c.data[DefaultSection] {
		options[i] = s
		i++
	}
	for s, _ := range c.data[section] {
		options[i] = s
		i++
	}

	return options, nil
}

// HasOption checks if the configuration has the given option in the section.
// It returns false if either the option or section do not exist.
func (c *Config) HasOption(section string, option string) bool {
	if section == "" {
		section = "default"
	}
	section = strings.ToLower(section)
	option = strings.ToLower(option)

	if _, ok := c.data[section]; !ok {
		return false
	}

	_, okd := c.data[DefaultSection][option]
	_, oknd := c.data[section][option]

	return okd || oknd
}

// RawString gets the (raw) string value for the given option in the section.
// The raw string value is not subjected to unfolding, which was illustrated in the beginning of this documentation.
// It returns an error if either the section or the option do not exist.
func (c *Config) RawString(section string, option string) (value string, err error) {
	if section == "" {
		section = "default"
	}

	section = strings.ToLower(section)
	option = strings.ToLower(option)

	if _, ok := c.data[section]; ok {
		if value, ok = c.data[section][option]; ok {
			return value, nil
		}
		return "", GetError{OptionNotFound, "", "", section, option}
	}
	return "", GetError{SectionNotFound, "", "", section, option}
}

// String gets the string value for the given option in the section.
// If the value needs to be unfolded (see e.g. %(host)s example in the beginning of this documentation),
// then String does this unfolding automatically, up to DepthValues number of iterations.
// It returns an error if either the section or the option do not exist, or the unfolding cycled.
func (c *Config) String(section string, option string) (value string, err error) {
	value, err = c.RawString(section, option)
	if err != nil {
		return "", err
	}

	return value, nil
}

// StringList gets the string values for the given option in the section.
// It returns an error if either the section or the option do not exist,
// or the unfolding cycled.
func (c *Config) StringList(section string, option string) (values []string, err error) {
	value, err := c.RawString(section, option)
	if err != nil {
		return nil, err
	}

	if strings.Contains(value, "\n") {
		v := strings.Split(value, "\n")
		if len(v[0]) == 0 {
			// Remove first empty element
			copy(v[0:], v[0+1:])
			v[len(v)-1] = ""
			v = v[:len(v)-1]
		}
		for _, val := range v {
			values = append(values, strings.Trim(val, " \t\r\n"))
		}
	} else {
		v := strings.NewReader(value)
		csvData := csv.NewReader(v)
		values, err = csvData.Read()
		if err != nil {
			return nil, err
		}
		for i, val := range values {
			values[i] = strings.Trim(val, " \t\r\n")
		}
	}

	return values, nil
}

// Int has the same behaviour as String but converts the response to int.
func (c *Config) Int(section string, option string) (value int, err error) {
	sv, err := c.String(section, option)
	if err == nil {
		value, err = strconv.Atoi(sv)
		if err != nil {
			err = GetError{CouldNotParse, "int", sv, section, option}
		}
	}

	return value, err
}

// IntList has the same behaviour as StringList but converts the response
// to []int.
func (c *Config) IntList(section string, option string) (values []int, err error) {
	slvs, err := c.StringList(section, option)
	if err != nil {
		return nil, err
	}

	for _, val := range slvs {
		value, err := strconv.Atoi(val)
		if err != nil {
			err = GetError{CouldNotParse, "int", val, section, option}
			return nil, err
		}
		values = append(values, value)
	}

	return values, nil
}

// Int64 has the same behaviour as String but converts the response
// to int64.
func (c *Config) Int64(section string, option string) (value int64, err error) {
	sv, err := c.String(section, option)
	if err == nil {
		value, err = strconv.ParseInt(sv, 10, 64)
		if err != nil {
			err = GetError{CouldNotParse, "int64", sv, section, option}
		}
	}

	return value, err
}

// Int64List has the same behaviour as StringList but converts the response
// to []int64.
func (c *Config) Int64List(section string, option string) (values []int64, err error) {
	slvs, err := c.StringList(section, option)
	if err != nil {
		return nil, err
	}

	for _, val := range slvs {
		value, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			err = GetError{CouldNotParse, "int64", val, section, option}
			return nil, err
		}
		values = append(values, value)
	}

	return values, nil
}

// Float64 has the same behaviour as String but converts the response
// to float64.
func (c *Config) Float64(section string, option string) (value float64, err error) {
	sv, err := c.String(section, option)
	if err == nil {
		value, err = strconv.ParseFloat(sv, 64)
		if err != nil {
			err = GetError{CouldNotParse, "float64", sv, section, option}
		}
	}

	return value, err
}

// Float64List has the same behaviour as StringList but converts the response
// to []float64.
func (c *Config) Float64List(section string, option string) (values []float64, err error) {
	slvs, err := c.StringList(section, option)
	if err != nil {
		return nil, err
	}

	for _, val := range slvs {
		value, err := strconv.ParseFloat(val, 64)
		if err != nil {
			print(err.Error())
			err = GetError{CouldNotParse, "float64", val, section, option}
			return nil, err
		}
		values = append(values, value)
	}

	return values, nil
}

// Bool has the same behaviour as String but converts the response to bool.
// See constant BoolStrings for string values converted to bool.
func (c *Config) Bool(section string, option string) (value bool, err error) {
	sv, err := c.String(section, option)
	if err != nil {
		return false, err
	}

	value, ok := BoolStrings[strings.ToLower(sv)]
	if !ok {
		return false, GetError{CouldNotParse, "bool", sv, section, option}
	}

	return value, nil
}

// BoolList has the same behaviour as StringList but converts the response
// to []bool.
func (c *Config) BoolList(section string, option string) (values []bool, err error) {
	slvs, err := c.StringList(section, option)
	if err != nil {
		return nil, err
	}

	for _, val := range slvs {
		value, ok := BoolStrings[strings.ToLower(val)]
		if !ok {
			err = GetError{CouldNotParse, "int64", val, section, option}
			return nil, err
		}
		values = append(values, value)
	}

	return values, nil
}
