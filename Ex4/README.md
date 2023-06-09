

Write a program which accepts a sequence of comma-separated numbers from console and 
generate an slice out of them. Return the slice.

Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.

Then, the output should be: [34 67 55 33 12 98]

## String is a data type in Go language and represents a slice of bytes that are UTF-8 encoded. This means that strings in Go language can represent a vast variety of characters. Strings are enclosed within double-quotes "". These strings are read-only and their values cannot be changed once defined.

### Split: This function splits a string into all substrings separated by the given separator and returns a slice that contains these substrings.
Syntax: 

func Split(str, sep string) []string

str is the string and sep is the separator.

### Trim: This function is used to trim the string all the leading and trailing Unicode code points which are specified in this function.

Syntax:

func Trim(str string, cutstr string) string

Here, str represent the current string and cutstr represents the elements which you want to trim in the given string.

e.g.

str1 := "!!Welcome to GeeksforGeeks !!"

 res1 := strings.Trim(str1, "!")

 fmt.Println("Result 1: ", res1) //Output: Result 1:  Welcome to GeeksforGeeks 


### How to iterate over a string?: You can iterate over string using for range loop. This loop can iterate over the Unicode code point for a string. Syntax:

for index, chr:= range str{
     // Statement..
}

Here, the index is the variable which store the first byte of UTF-8 encoded code point and chr store the characters of the given string and str is a string