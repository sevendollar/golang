package jef

import "fmt"

// SliceOfCharacters ...
// str <- string
// index <- first index argument represents the beginning index
// index <- second (last) index argument represents the ending index
func SliceOfCharacters(str string, index ...int) string {
	// to store the index value of "str" argument
	// ss <- slice of string
	var ss []int

	// to store the value of "index" argument
	// b <- begin index
	// e <- end index
	var b, e int

	// find the range of character
	for i := range str {
		ss = append(ss, i)
	}

	// version 1 using if statement
	//
	// if len(index) > 1 { // two arguments of index
	// 	b = ss[index[0]]

	// 	// the last character
	// 	if len(ss)-index[1] == 0 {
	// 		return str[b:len(str)]
	// 	}
	// 	e = ss[index[1]]
	// 	return str[b:e]

	// } else if len(index) == 1 { // one argument of index
	// 	b = ss[index[0]]

	// 	// the last character
	// 	if len(ss)-index[0] == 1 {
	// 		return str[b:len(str)]
	// 	}
	// 	e = ss[index[0]+1]
	// 	return str[b:e]
	// }
	// fmt.Print("error: not enough arguments")
	// return ""

	// version 2 using switch statement
	//
	// slice the string base on indexes
	switch l := len(index); {
	case l > 1: // two arguments of index
		b = ss[index[0]]

		// the last character
		if len(ss)-index[1] == 0 {
			return str[b:len(str)]
		}
		e = ss[index[1]]
		return str[b:e]

	case l == 1: // one argument of index
		b = ss[index[0]]

		// the last character
		if len(ss)-index[0] == 1 {
			return str[b:len(str)]
		}
		e = ss[index[0]+1]
		return str[b:e]

	default:
		fmt.Print("error: not enough arguments")
		return ""
	}

}
