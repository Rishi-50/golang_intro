package main

import "fmt"

func main()  {
	// var str1 string = "GO"
	// var str2 string = "strings"

	// fmt.Println(str1, str2)

	// var str string = "GO"
	// var exint int = 10
	// var exfloat float64 = 43.3

	// fmt.Println(str, exint, exfloat)


	// var(
	// 	empid int = 204
	// 	fname , lname string = "Rishi", "Savla" 
	// )
	// fmt.Println(fname,lname,empid)

	// name := "John Doe"
	// empid , age, isDeveloper := 201 , 34, true				// ":="  --> only works while declaring a variable in a function (not outside of a function)

	// fmt.Println(name,empid,age,isDeveloper)


	var(
		fname,lname string
		age int 
		salary float64
		isDeveloper bool
	)

	fmt.Printf("First Name: %s, Last Name: %s, Age: %d, Salary:%f, Is Developer: %t\n" , fname, lname, age,salary,isDeveloper)
}