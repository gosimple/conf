/*
This package implements a parser for configuration files.
This allows easy reading and writing of structured configuration files.

Given the configuration file:

	[default]
	host = example.com
	port = 443
	php = on

	[service-1]
	host = s1.example.com
	allow-writing = false

To read this configuration file, do:

	c, err := conf.ReadFile("server.conf")
	c.String("default", "host")             // returns example.com
	c.Int("", "port")                       // returns 443 (assumes "default")
	c.Bool("", "php")                       // returns true
	c.String("service-1", "host")           // returns s1.example.com
	c.Bool("service-1","allow-writing")     // returns false
	c.Int("service-1", "port")              // returns 0 and a GetError

Note that all section and option names are case insensitive. All values
are case sensitive.

Goconfig's string substitution syntax has not been removed. However, it may be
taken out or modified in the future.
*/
package conf
