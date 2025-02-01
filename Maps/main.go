package main

import (
	"fmt"
	"maps"
)

func main()  {
	
	m := make(map[string]int, 0)     // make(map[key-type]val-type)

	m["key1"] = 10
	m["key2"] = 20
	m["key3"] = 30
 
	fmt.Println(m)

	v1 := m["key1"]
	v2 := m["key2"]
	v3 := m["key3"]

	fmt.Println(v1,v2,v3)
	_, prs := m["key2"]
    fmt.Println("prs:", prs)

	n1 := map[string]int{"foo": 1, "bar": 2}
	 
	n2 := map[string]int{"foo": 2, "bar": 2}

	if maps.Equal(n1, n2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

}