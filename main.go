package main

import "fmt"

func main() {
	x:=10
	fmt.Printf("Value of x is %d\n",x)
	fmt.Printf("Address of x = %x\n",&x)

	var p *int
	p=&x	//pointer p get address var x
	fmt.Printf("Value of p = %x\n",p)
	fmt.Printf("Address of p = %x\n",&p)


	*p=20	//pointer p change value of x
	fmt.Printf("Value of x = %d\n",x) //after change value by pointer
}

