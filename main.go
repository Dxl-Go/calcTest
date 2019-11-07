package main

import (
	"fmt"
	"gitTest/calc"
)

func main()  {
	fmt.Println("calc.Add called!")
	res:=calc.Add(10,20)
	fmt.Println("res:",res)
	fmt.Println("calc.Sub called!")
	res=calc.Sub(10,20)
	fmt.Println("res:",res)
	fmt.Println("calc.Multi called!")
	res=calc.Multi(10,20)
	fmt.Println("10*20=:",res)
}
