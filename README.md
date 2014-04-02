conf
====

Package `conf` provide support for parsing configuration files.

[Documentation online](http://godoc.org/github.com/gosimple/conf)

## Installation

	go get -u github.com/gosimple/conf

## Usage

Check `example` folder

	import "github.com/gosimple/conf"

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

<https://github.com/gosimple/conf/issues>

### Info

Package conf is based on [goconfig](https://code.google.com/p/goconf/) 

## License

The source files are distributed under the The BSD 3-Clause License.
You can find full license text in the `LICENSE` file.
