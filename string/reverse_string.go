package reverse_string // name of package to be called, ie "reverse_string.Reverse("stuff")

func Reverse(s string) string { // name of funtion, type of parameter passed in which
                                // will be referenced in the function as 's',
                                // return type then opening brace
	b := []rune(s) // the bytes of the string- will break on some unicode chars
                  // the := is shorthand for assigning a variable without an explicit type
                  // 'age := 26' is the same as 'var age int = 26'
	for i := 0; i < len(b)/2; i++ { // for loop, starts with the variable initialization,
                                  // then the condition statement: when i < 1/2 oflength of string
                                  // and the incrementation of the variable
		j := len(b)-i-1 // black magic??? length of b minus index minus one
		b[i], b[j] = b[j], b[i] // reverse
	}
	return string(b) //return with type specified
}
