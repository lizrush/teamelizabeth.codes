package reverse_string // name of package to be called, ie "reverse_string.Reverse("stuff")

func Reverse(s string) string { // name of funtion, type of thing passed in which
                                // will be referenced in the function as 's',
                                // then i think the type of the output?
	b := []byte(s) // the bytes of the string- will break on some unicode chars
	for i := 0; i < len(b)/2; i++ { // for loop sorta like js, with incrementor,
                                  // when i < 1/2 oflength of string in bytes
		j := len(b)-i-1 // black magic???
		b[i], b[j] = b[j], b[i] // reverse
	}
	return string(b) //return with type specified
}
