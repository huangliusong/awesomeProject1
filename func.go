package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a,b int ,op string) (int,error){
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		q,_:=div(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf("unsupported operation:%s"+op)
	}
}

func div(a,b int ,)(q int, r int){
	q= a/b
	r=a%b
	return q,r
}

func apply(op func(int,int)int ,a,b int )int{
	p:=reflect.ValueOf(op).Pointer()
	opName:=runtime.FuncForPC(p).Name()
	fmt.Println("Calling func %s"+"(%d,%d)",opName,a,b)
	return op(a,b)
}

func pow(a,b int) int{
	return int(math.Pow(float64(a),float64(b)))
}
func main() {
	if result,err:=eval(3,4,"/");err!=nil{
		fmt.Println("error",err)
	}else{
		fmt.Println(result)
	}
	fmt.Println(eval(3,4,"/"))
	fmt.Println(div(13,3))
	q,r:=div(13,3)
	fmt.Println(q,r)
	fmt.Println(apply(func(a int,b int) int{
		return int(math.Pow(float64(a),float64(b)))
	},3,4))
}
