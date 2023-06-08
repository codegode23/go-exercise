# Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.

// An Empty map
map[Key_Type]Value_Type{}

// Map with key-value pair
map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

## make 
is a special function in Go that can take a different number of types and arguments.

It returns an instance of the type passed as the first argument:

### syntax

obj := make(someType, optionalArgument1, optionalArgument2)

We can initialize slices of any type using make:

words := make([]string, 2)

Here, the first argument is the type and the second argument is the length.

By default, a new slice is initialized and filled with as many empty values as the length specified.

So, in this case, the value of words would be []string{"", ""}
