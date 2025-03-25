package main

import "fmt"

type Display interface {
	update() string
}

type show struct {
	value string
}

func (s show) update() string {
	return s.value
}

func main() {
	var d Display
	d = show{"NSB from TYBCA"}
	if val, ok := d.(show); ok {
		fmt.Println(val.value)
	} else {
		fmt.Println("Value is not a String")
	}
}
