package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {

	sentence := "This is the tutorial for golang"
	
    fmt.Println(strings.Contains(sentence,"golang"))
    fmt.Println(strings.ReplaceAll(sentence,"golang","GO"))
    fmt.Println(strings.Count(sentence,"golang"))
    fmt.Println(strings.ToUpper(sentence))
    fmt.Println(strings.Split(sentence," "))

    ages := []int{10,4,2,65,30}

    sort.Ints(ages) 
    fmt.Println(ages)
    
    f := sort.SearchInts(ages,30)
    fmt.Println(f)

    names := []string{"rishi","ruchit","neel","jinay"}

    sort.Strings(names)
    fmt.Println(names)

    s := sort.SearchStrings(names,"rishi")
    fmt.Println(s)
}