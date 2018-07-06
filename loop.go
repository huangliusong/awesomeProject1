package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func converToBin(n int)string{
	result:=""
	for ;n>0 ; n/=2 {
		lsb:=n%2
		result=strconv.Itoa(lsb)+result
	}
	return  result
}
 func printFile(filename string){
	 file,err:=	os.Open(filename)
	 if err!=nil{
		panic(err)
	 }
	 scanner:=bufio.NewScanner(file)
	 for scanner.Scan(){
		 fmt.Println(scanner.Text())
	 }
 }

 func forever(){
 	for{
 		fmt.Println("abc")
	}
 }
func main() {
	fmt.Println(
		converToBin(5),
		converToBin(13),
		converToBin(72387885),
		converToBin(0))

	printFile("abc.txt")
	forever()
}
