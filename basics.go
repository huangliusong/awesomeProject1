package main

import (
	"fmt"
	"math"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a,s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDedcction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func consts(){
	const(
		filename="abc.txt"
		a,b=3,4
	)
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(filename,c)
}

func triangle(){
	var a,b int =3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}

func enums(){
	const(
		cpp=iota
		_
		python
		golang
		javascript
	)
	const(
		b=1<<(10*iota)
		kb
		mb
		tb
		pb
	)
	fmt.Println(b,kb,mb,tb,pb)
	fmt.Println(cpp,javascript,python,golang)
}

func bounded(v int) int {
	if v>100{

	}
}
func main() {
	fmt.Println("你好啊?");
	variableZeroValue();
	variableInitialValue();
	variableTypeDedcction();
	variableShorter();
	fmt.Println(aa,bb,ss)
	triangle();
	consts();
	enums();
}
