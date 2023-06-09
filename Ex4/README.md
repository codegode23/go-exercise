

Write a program which accepts a sequence of comma-separated numbers from console and generate an slice out of them. Return the slice.

Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.

Then, the output should be: [34 67 55 33 12 98]

## String is a data type in Go language and represents a slice of bytes that are UTF-8 encoded. This means that strings in Go language can represent a vast variety of characters. Strings are enclosed within double-quotes "". These strings are read-only and their values cannot be changed once defined.

How to iterate over a string?: You can iterate over string using for range loop. This loop can iterate over the Unicode code point for a string. Syntax:

for index, chr:= range str{
     // Statement..
}

Here, the index is the variable which store the first byte of UTF-8 encoded code point and chr store the characters of the given string and str is a string