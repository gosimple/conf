// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"bitbucket.org/gosimple/conf"
)

func main() {
	c, _ := conf.ReadFile("config.ini")

	host, _ := c.String("default", "host")
	fmt.Println(host)

	port, _ := c.Int("", "port")
	fmt.Println(port)

	php, _ := c.Bool("", "php")
	fmt.Println(php)

	serHost, _ := c.String("service-1", "host")
	fmt.Println(serHost)

	serAW, _ := c.Bool("service-1", "allow-writing")
	fmt.Println(serAW)

	serPort, err := c.Int("service-1", "port")
	if err != nil {
		fmt.Println(serPort, "Error:", err)
	} else {
		fmt.Println(serPort)
	}

	myList, _ := c.StringList("list", "list-1")
	fmt.Println(myList)
	myList2, _ := c.StringList("list", "list-2")
	fmt.Println(myList2)
	myList3, _ := c.StringList("list", "list-3")
	fmt.Println(myList3)

	myList4, _ := c.IntList("list", "list-4")
	fmt.Println(myList4)
	myList5, _ := c.Int64List("list", "list-4")
	fmt.Println(myList5)
	myList6, _ := c.Float64List("list", "list-5")
	fmt.Println(myList6)

	myList7, _ := c.BoolList("list", "list-6")
	fmt.Println(myList7)
}
