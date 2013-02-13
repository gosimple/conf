// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package conf

import (
	"fmt"
	"regexp"
	"strings"
)

// Config is the representation of configuration settings.
// The public interface is entirely through methods.
type Config struct {
	data map[string]map[string]string // Maps sections to options to values.
}

const (
	// Get Errors
	SectionNotFound = iota
	OptionNotFound
	MaxDepthReached

	// Read Errors
	BlankSection

	// Get and Read Errors
	CouldNotParse
)

var (
	DefaultSection = "default" // Default section name (must be lower-case).
	DepthValues    = 200       // Maximum allowed depth when recursively substituing variable names.

	// Strings accepted as bool.
	BoolStrings = map[string]bool{
		"t":     true,
		"true":  true,
		"y":     true,
		"yes":   true,
		"on":    true,
		"1":     true,
		"f":     false,
		"false": false,
		"n":     false,
		"no":    false,
		"off":   false,
		"0":     false,
	}

	varRegExp = regexp.MustCompile(`%\(([a-zA-Z0-9_.\-]+)\)s`)
)

// AddSection adds a new section to the configuration.
// It returns true if the new section was inserted, and false if the section already existed.
func (c *Config) AddSection(section string) bool {
	section = strings.ToLower(section)

	if _, ok := c.data[section]; ok {
		return false
	}
	c.data[section] = make(map[string]string)

	return true
}

// RemoveSection removes a section from the configuration.
// It returns true if the section was removed, and false if section did not exist.
func (c *Config) RemoveSection(section string) bool {
	section = strings.ToLower(section)

	switch _, ok := c.data[section]; {
	case !ok:
		return false
	case section == DefaultSection:
		return false // default section cannot be removed
	default:
		for o, _ := range c.data[section] {
			delete(c.data[section], o)
		}
		delete(c.data, section)
	}

	return true
}

// AddOption adds a new option and value to the configuration.
// It returns true if the option and value were inserted, and false if the value was overwritten.
// If the section does not exist in advance, it is created.
func (c *Config) AddOption(section string, option string, value string) bool {
	c.AddSection(section) // make sure section exists

	section = strings.ToLower(section)
	option = strings.ToLower(option)

	_, ok := c.data[section][option]
	c.data[section][option] = value

	return !ok
}

// RemoveOption removes a option and value from the configuration.
// It returns true if the option and value were removed, and false otherwise,
// including if the section did not exist.
func (c *Config) RemoveOption(section string, option string) bool {
	section = strings.ToLower(section)
	option = strings.ToLower(option)

	if _, ok := c.data[section]; !ok {
		return false
	}

	_, ok := c.data[section][option]
	delete(c.data[section], option)

	return ok
}

// New creates an empty configuration representation.
// This representation can be filled with AddSection and AddOption and then
// saved to a file using WriteFile.
func New() *Config {
	c := new(Config)
	c.data = make(map[string]map[string]string)

	c.AddSection(DefaultSection) // default section always exists

	return c
}

type GetError struct {
	Reason    int
	ValueType string
	Value     string
	Section   string
	Option    string
}

func (err GetError) Error() string {
	switch err.Reason {
	case SectionNotFound:
		return fmt.Sprintf("section '%s' not found", string(err.Section))
	case OptionNotFound:
		return fmt.Sprintf("option '%s' not found in section '%s'", string(err.Option), string(err.Section))
	case CouldNotParse:
		return fmt.Sprintf("could not parse %s value '%s'", string(err.ValueType), string(err.Value))
	case MaxDepthReached:
		return fmt.Sprintf("possible cycle while unfolding variables: max depth of %d reached", int(DepthValues))
	}

	return "invalid get error"
}

type ReadError struct {
	Reason int
	Line   string
}

func (err ReadError) Error() string {
	switch err.Reason {
	case BlankSection:
		return "empty section name not allowed"
	case CouldNotParse:
		return fmt.Sprintf("could not parse line: %s", string(err.Line))
	}

	return "invalid read error"
}
