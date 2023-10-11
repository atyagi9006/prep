package main

import "fmt"

func main() {
	isValid("amit")
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := []string{}
	open := make(map[string]string)
	open["("] = ")"
	open["{"] = "}"
	open["["] = "]"
	fmt.Printf("%T %v \n", string(s[0]), string(s[0]))
	for _, r := range s {
		if closed, ok := open[string(r)]; ok {
			stack = append(stack, closed)
			continue
		}


		l := len(stack) - 1
		if l < 0 && string(r) != stack[l] {
			return false
		}
		stack = stack[:l]
	}

	return len(stack) == 0
}

func pop(stack[]string)string{
	l:=len(stack)-1
	last:=stack[l]
	stack= stack[:l]
	return last
}
