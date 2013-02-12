package conf

import (
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

// Int has the same behaviour as GetString but converts the response to int.
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

// Int64 has the same behaviour as GetString but converts the response to int64.
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

// Float has the same behaviour as GetString but converts the response to float.
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

// Bool has the same behaviour as GetString but converts the response to bool.
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
