conf
====

Package `conf` provide support for parsing configuration files.

[Documentation online](http://godoc.org/bitbucket.org/gosimple/conf)

## Installation

	go get bitbucket.org/gosimple/conf

## Example

Check `example` folder

	import "bitbucket.org/gosimple/conf"

NOTE: All section names and options are case insensitive. All values are case
sensitive.

### Example 1

**Config**

	host = something.com
	port = 443
	active = true
	compression = off
	
	list-str = hello, world
	list-int = 1, 2, 3

**Code**

	c, err := conf.ReadFile("something.config")
	c.String("default", "host")				// return something.com
	c.Int("default", "port")				// return 443
	c.Bool("default", "active")				// return true
	c.Bool("default", "compression")		// return false
	
	c.StringList("default", "list-str")		// return ["hello", "world"]
	c.IntList("default", "list-int")		// return [1, 2, 3]

### Example 2

**Config**

	[default]
	host = something.com
	port = 443
	active = true
	compression = off
	
	[service-1]
	compression = on
	
	[service-2]
	port = 444

**Code**

	c, err := conf.ReadFile("something.config")
	c.Bool("default", "compression") // returns false
	c.Bool("service-1", "compression") // returns true
	c.Bool("service-2", "compression") // returns GetError

### Requests or bugs?

<https://bitbucket.org/gosimple/conf/issues>

### Info

Package conf is based on [goconfig](https://code.google.com/p/goconf/) 
