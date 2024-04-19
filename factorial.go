package main

import "fmt"

func factorial(n int,ch chan int){
	res,i:=1,1
	
	for i<=n{
		res*=i
		i++
	}
	ch<-res

}

func main()  {
	n:=0
	chan1:=make(chan int,1)
	fmt.Print("Enter the number: ")
	fmt.Scanln(&n)
	go factorial(n,chan1)
	fmt.Printf("Factorial of %v = %v",n,<-chan1)

	
}